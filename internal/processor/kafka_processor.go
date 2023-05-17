package processor

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

// MessageProcessor is the interface for processing messages
type MessageProcessor interface {
	StartProcessing() error
}

type KafkaProcessor struct {

	// Host is the host of the kafka broker
	Host string

	// Topic is the topic to consume
	Topic string

	// ConsumerGroup is the consumer group to use
	ConsumerGroup string

	// Reader is the kafka reader
	reader *kafkago.Reader

	// Writer is the kafka writer
	writer *kafkago.Writer

	tlsConfig *tls.Config

	// Tracer is the opentelemetry tracer
	tracer trace.Tracer
}

func (p *KafkaProcessor) Init() error {
	slog.Info(fmt.Sprintf("connect to kafka host %s", p.Host))

	p.tracer = otel.Tracer("kafka-processor")

	clientCertPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(utils.KafkaCA())

	if err != nil {
		slog.Error("error reading ca cert", err)
		return err
	}

	if ok := clientCertPool.AppendCertsFromPEM(ca); !ok {
		slog.Warn("error appending certs from pem")
	}

	mechanism, err := scram.Mechanism(scram.SHA256, utils.KafkaUser(), utils.KafkaPassword())
	if err != nil {
		slog.Error("error creating scram mechanism", err)
		return err
	}

	cert, _ := tls.LoadX509KeyPair(utils.KafkaCert(), utils.KafkaKey())

	dialer := &kafkago.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS: &tls.Config{
			ClientCAs:          clientCertPool,
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{cert},
		},
	}

	p.reader = kafkago.NewReader(kafkago.ReaderConfig{
		Brokers:  []string{p.Host},
		Topic:    p.Topic,
		GroupID:  p.ConsumerGroup,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Dialer:   dialer,
	})

	p.writer = kafkago.NewWriter(kafkago.WriterConfig{
		Brokers: []string{p.Host},
		Topic:   p.Topic,
		Dialer:  dialer,
	})

	return p.validateSettings()
}

func (p *KafkaProcessor) WriteMessage(ctx context.Context, msg kafkago.Message) error {
	ctx, span := p.tracer.Start(ctx, "write-message")
	defer span.End()

	span.SetAttributes(attribute.String("key", string(msg.Key)))
	span.SetAttributes(attribute.String("topic", p.Topic))

	return p.writer.WriteMessages(ctx, msg)
}

func (p *KafkaProcessor) Close() error {
	return p.reader.Close()
}

func (p *KafkaProcessor) validateSettings() error {
	if p.Host == "" {
		return errors.New("host not set")
	}

	if p.Topic == "" {
		return errors.New("topic not set")
	}

	if p.ConsumerGroup == "" {
		return errors.New("consumer group not set")
	}

	return nil
}

func (p *KafkaProcessor) CommitMessage(ctx context.Context, msg *kafkago.Message) error {
	ctx, span := p.tracer.Start(ctx, "commit-message")
	defer span.End()
	span.SetAttributes(attribute.String("key", string(msg.Key)))
	span.SetAttributes(attribute.String("topic", p.Topic))

	err := p.reader.CommitMessages(ctx, *msg)
	if err != nil {
		slog.Error("error committing message", err)
		span.RecordError(err)
		return err
	}
	return nil
}

// CollectMessages collects messages from the kafka reader
// and returns them in a channel
// if the kafka read returns an error the function panics
// The caller is responsible for committing the messages
func (p *KafkaProcessor) CollectMessages(ctx context.Context, channelLength int) <-chan *kafkago.Message {
	result := make(chan *kafkago.Message, channelLength)

	slog.Info(fmt.Sprintf("start processing messages from %s with cg %s", p.Topic, p.ConsumerGroup))
	go func() {
		defer close(result)
		for {
			msg, err := p.reader.FetchMessage(ctx)

			if err == io.EOF {
				slog.Info(fmt.Sprintf("reader for topic %s closed", p.Topic))
				time.Sleep(5 * time.Second)
				continue
			}

			if err != nil {
				slog.Error("error reading message", err)
				time.Sleep(5 * time.Second)
				continue
			}

			result <- &msg
		}
	}()

	return result
}
