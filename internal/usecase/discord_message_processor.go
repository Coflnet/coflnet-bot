package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	"github.com/rs/zerolog/log"
	kafkago "github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

const (
    discordMessageTracerName = "discord-message-processor"
)

type DiscordMessageProcessor struct {
	kafkaProcessor *KafkaProcessor
    tracer trace.Tracer
    discordHandler *discord.DiscordHandler
}

func (p *DiscordMessageProcessor) StartProcessing() error {
	p.kafkaProcessor = &KafkaProcessor{
		Host:          utils.KafkaHost(),
		Topic:         utils.FlipSummaryTopic(),
		ConsumerGroup: utils.FlipSummaryConsumerGrup(),
	}
    p.tracer = otel.Tracer(flipProcessorTracerName)

    var err error
    p.discordHandler, err = discord.NewDiscordHandler()
    if err != nil {
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

	for msg := range msgs {
		go func(msg *kafkago.Message) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

            _, span := otel.Tracer(flipProcessorTracerName).Start(ctx, "process-discord-message-kafka-message")
            defer span.End()

			err := p.processMessage(ctx, msg)
			if err != nil {
				log.Error().Err(err).Msg("error processing flip summary")
                span.RecordError(err)
				return
			}
            span.AddEvent("processed message")

			p.kafkaProcessor.CommitMessage(ctx, msg)
            span.AddEvent("committed message")
		}(msg)
	}

	slog.Warn(fmt.Sprintf("kafka processor for topic %s stopped", p.kafkaProcessor.Topic))
	return nil
}

func (p *DiscordMessageProcessor) processMessage(ctx context.Context, msg *kafkago.Message) error {
    _, span := otel.Tracer(flipProcessorTracerName).Start(ctx, "process-discord-message")
    defer span.End()

    // deserialize message
	var m *discord.DiscordMessage
    err := json.Unmarshal(msg.Value, msg)
	if err != nil {
        slog.Error("error unmarshalling message %s", err)
        span.RecordError(err)
		return err
	}

    // call discord
    err = p.discordHandler.SendMessage(ctx, m)
    if err != nil {
        slog.Error("error sending message to discord %s", err)
        span.RecordError(err)
        return err
    }
    
    slog.Debug("processed discord kafka message %s", msg.Key)
	return nil
}