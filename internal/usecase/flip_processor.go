package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/rs/zerolog/log"
	kafkago "github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

const (
    flipProcessorTracerName = "flip-processor"
)

type FlipProcessor struct {
	kafkaProcessor *KafkaProcessor
    tracer trace.Tracer
}

func (p *FlipProcessor) StartProcessing() error {
	p.kafkaProcessor = &KafkaProcessor{
		Host:          utils.KafkaHost(),
		Topic:         utils.FlipSummaryTopic(),
		ConsumerGroup: utils.FlipSummaryConsumerGrup(),
	}
    p.tracer = otel.Tracer(flipProcessorTracerName)

	err := p.kafkaProcessor.Init()
	if err != nil {
		return err
	}
	defer p.kafkaProcessor.Close()

	ctx := context.Background()
	msgs := p.kafkaProcessor.CollectMessages(ctx, 100)
	semaphore := make(chan struct{}, 10)

    slog.Info("starting to process flip summaries")
	for msg := range msgs {
		semaphore <- struct{}{}
		go func(msg *kafkago.Message) {
			defer func() { <-semaphore }()

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()


            ctx, span := otel.Tracer(flipProcessorTracerName).Start(ctx, "process-flipsummary-kafka-message")
            defer span.End()

			err := p.processMessage(ctx, msg)
			if err != nil {
				slog.Error("error processing flip summary", err)
                span.RecordError(err)
				return
			}
            span.AddEvent("processed message")

			err = p.kafkaProcessor.CommitMessage(ctx, msg)
            if err != nil {
                slog.Error("error committing message", err)
                span.RecordError(err)
                return
            }

            span.AddEvent("committed message")
		}(msg)
	}

	log.Panic().Msgf("kafka processor for topic %s stopped", p.kafkaProcessor.Topic)
	return nil
}

func (p *FlipProcessor) processMessage(ctx context.Context, msg *kafkago.Message) error {
    ctx, span := otel.Tracer(flipProcessorTracerName).Start(ctx, "process-flipsummary")
    defer span.End()


    var flip model.Flip
    err := json.Unmarshal(msg.Value, &flip)
    if err != nil {
        return err
    }
    span.SetAttributes(attribute.Int("profit", flip.Profit))

    if flip.Profit > 10_000_000 {
        slog.Info(fmt.Sprintf("flip with profit %d", flip.Profit))
        return nil
    }


	return nil
}
