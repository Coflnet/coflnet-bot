package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/Coflnet/coflnet-bot/internal/utils"
    pubdiscord "github.com/Coflnet/coflnet-bot/pkg/discord"
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
    discordHandler *discord.DiscordHandler
}

func (p *FlipProcessor) StartProcessing() error {
	p.kafkaProcessor = &KafkaProcessor{
		Host:          utils.KafkaHost(),
		Topic:         utils.FlipSummaryTopic(),
		ConsumerGroup: utils.FlipSummaryConsumerGrup(),
	}
    p.tracer = otel.Tracer(flipProcessorTracerName)

    var err error
    p.discordHandler, err = discord.NewDiscordHandler()
    if err != nil {
        slog.Error("error initializing discord handler", err)
        return err
    }

	err = p.kafkaProcessor.Init()
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

	return errors.New("flip processor kafka channel closed")
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
    span.SetAttributes(attribute.String("end", flip.Sell.End.Format(time.RFC3339)))


    // if sell end is older than 24 hours ignore
    if flip.Sell.End.Unix() < time.Now().Add(-24*time.Hour).Unix() {
        span.AddEvent("skip-flip-because-it-is-old")
        slog.Info(fmt.Sprintf("skip flip because it is old"))
        return nil
    }

    if flip.Profit > 10_000_000 {

        err = p.discordHandler.SendMessage(ctx, &discord.DiscordMessage{
            Message: fmt.Sprintf("flip with profit %d", flip.Profit),
            Channel: pubdiscord.FlipChannel,
        })

        slog.Info(fmt.Sprintf("flip with profit %d", flip.Profit))
        return nil
    }


	return nil
}
