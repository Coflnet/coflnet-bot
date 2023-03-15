package usecase

import (
	"context"
	"errors"

	kafkago "github.com/segmentio/kafka-go"
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
}

func (p *KafkaProcessor) Init() error {
    p.reader = kafkago.NewReader(kafkago.ReaderConfig{
        Brokers:   []string{p.Host},
        Topic:     p.Topic,
        GroupID:   p.ConsumerGroup,
        MinBytes:  10e3, // 10KB
        MaxBytes:  10e6, // 10MB
    })

    return p.validateSettings()
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
    return p.reader.CommitMessages(ctx, *msg)
}

// CollectMessages collects messages from the kafka reader
// and returns them in a channel
// if the kafka read returns an error the function panics
// The caller is responsible for committing the messages
func (p *KafkaProcessor) CollectMessages(ctx context.Context, channelLength int) <-chan *kafkago.Message {
    result := make(chan *kafkago.Message, channelLength)
    go func() {
        defer close(result)
        for {
            msg, err := p.reader.FetchMessage(ctx)

            if err != nil {
                slog.Error("error reading message", err)
                continue
            }

            result <- &msg
        }
    }()
    return result
}

