package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/segmentio/kafka-go"
	"github.com/vmihailenco/msgpack"
	"go.opentelemetry.io/otel"
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

    return errors.New("kafka message channel closed")
}

func (p *McVerifyProcessor) processMessage(ctx context.Context, msg *kafka.Message) error {
    _, span := p.tracer.Start(ctx, "process-mc-verify-message")
    defer span.End()


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

    span.AddEvent(fmt.Sprintf("refreshed user with id %s", id))

    return nil
}


type McVerifyMessageContent struct {
	UserId   string `msgpack:"userId"`
	UUID     string `msgpack:"uuid"`
	Existing int    `msgpack:"existing"`
}
