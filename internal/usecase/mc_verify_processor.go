package usecase

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/segmentio/kafka-go"
	"github.com/vmihailenco/msgpack"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

const (
    mcVerifyProcessorName = "chat-processor"
)

type McVerifyProcessor struct {
    tracer trace.Tracer
    kafkaProcessor *KafkaProcessor
    userHandler *UserHandler
}

func (p *McVerifyProcessor) init() error {
    p.tracer = otel.Tracer(flipProcessorTracerName)
    var err error
    p.userHandler, err = NewUserHandler()

    if err != nil {
        slog.Error("failed to init user handler", err)
        return err
    }

    return nil
}

func (p *McVerifyProcessor) StartProcessing() error {
   if err := p.init(); err != nil { 
       return err
    }

    p.kafkaProcessor = &KafkaProcessor{
        Host:          utils.KafkaHost(),
        Topic:         utils.McVerifyKafkaTopic(),
        ConsumerGroup: utils.McVerifyKafkaConsumerGroup(),
    }

    err := p.kafkaProcessor.Init()
    if err != nil {
        return err
    }
    defer p.kafkaProcessor.Close()


    ctx := context.Background()
    slog.Info("starting mc verify processor")

    msgs := p.kafkaProcessor.CollectMessages(ctx, 100)
    semaphore := make(chan struct{}, 10)

    go func() {
        for msg := range msgs {
            go func(msg *kafka.Message) {
	    		semaphore <- struct{}{}
                defer func() { <-semaphore }()

                ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
                defer cancel()

                _, span := p.tracer.Start(ctx, "process-mc-verify-kafka-message")
                defer span.End()

                err := p.processMessage(ctx, msg)
                if err != nil {
                    slog.Error("failed to process message", err)
                    span.RecordError(err)
                    return
                }

                span.AddEvent("processed message")

                err = p.kafkaProcessor.CommitMessage(ctx, msg)
                if err != nil {
                    slog.Error("failed to commit message", err)
                    span.RecordError(err)
                    return
                }

                slog.Debug("mc verify message processed")
                span.AddEvent("comitted message")
            }(msg)
        }
        panic("kafka message channel closed")
    }()
    return nil
}

func (p *McVerifyProcessor) processMessage(ctx context.Context, msg *kafka.Message) error {
    ctx, span := p.tracer.Start(ctx, "process-mc-verify-message")
    defer span.End()
    span.SetAttributes(attribute.Int64("offset", msg.Offset))
    span.SetAttributes(attribute.String("key", string(msg.Key)))


    var verificationMessage McVerifyMessageContent
	err := msgpack.Unmarshal(msg.Value, &verificationMessage)
	if err != nil {
		slog.Error("failed to unmarshal mc_verify message", err)
        span.RecordError(err)
		return err
	}

	id, err := strconv.Atoi(verificationMessage.UserId)
	if err != nil {
        slog.Error("failed to convert user id to int", err)
        span.RecordError(err)
		return err
	}

    err = p.userHandler.RefreshUserId(ctx, id)
    if err != nil {
        slog.Error("failed to refresh user id", err)
        span.RecordError(err)
    }

    return nil
}


type McVerifyMessageContent struct {
	UserId   string `msgpack:"userId"`
	UUID     string `msgpack:"uuid"`
	Existing int    `msgpack:"existing"`
}
