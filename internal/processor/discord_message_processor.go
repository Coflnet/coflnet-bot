package processor

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/Coflnet/coflnet-bot/internal/discord"
	"github.com/Coflnet/coflnet-bot/internal/utils"
	kafkago "github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	discordMessageTracerName = "discord-message-processor"
)

func NewDiscordMessageProcessor(d *discord.DiscordHandler) *DiscordMessageProcessor {

	dp := &DiscordMessageProcessor{
		discordHandler: d,
		tracer:         otel.Tracer(discordMessageTracerName),
	}

	kafkaUrl, err := utils.KafkaHost()
	if err != nil {
		slog.Error("error getting kafka url", slog.String("err", err.Error()))
		return nil
	}

	kafkaProcessor := &KafkaProcessor{
		Host:          kafkaUrl,
		Topic:         utils.DiscordMessageTopic(),
		ConsumerGroup: utils.DiscordMessageConsumerGroup(),
	}
	dp.kafkaProcessor = kafkaProcessor

	return dp
}

// sends messages from a kafka topic to discord
type DiscordMessageProcessor struct {
	kafkaProcessor *KafkaProcessor
	tracer         trace.Tracer
	discordHandler *discord.DiscordHandler
}

func (p *DiscordMessageProcessor) StartProcessing() error {
	err := p.kafkaProcessor.Init()
	if err != nil {
		return err
	}

	ctx := context.Background()
	msgs := p.kafkaProcessor.CollectMessages(ctx, 100)
	semaphore := make(chan struct{}, 10)

	go func() {
		defer func(kafkaProcessor *KafkaProcessor) {
			err = kafkaProcessor.Close()
			if err != nil {
				slog.Error("error closing kafka processor", err)
			}
		}(p.kafkaProcessor)

		for msg := range msgs {
			semaphore <- struct{}{}
			go func(msg *kafkago.Message) {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()

				defer func() { <-semaphore }()

				_, span := p.tracer.Start(ctx, "process-discord-message-kafka-message")
				defer span.End()

				err = p.processMessage(ctx, msg)
				if err != nil {
					slog.Error("error processing message in topic "+p.kafkaProcessor.Topic, err)
					span.RecordError(err)
					return
				}
				span.AddEvent("processed message")

				err = p.kafkaProcessor.CommitMessage(ctx, msg)
				if err != nil {
					slog.Error("error committing message in topic "+p.kafkaProcessor.Topic, err)
					span.RecordError(err)
					return
				}

				span.AddEvent("committed message")
			}(msg)
		}
		panic("channel closed")
	}()
	return nil
}

func (p *DiscordMessageProcessor) processMessage(ctx context.Context, msg *kafkago.Message) error {
	_, span := p.tracer.Start(ctx, "process-discord-message")
	defer span.End()

	// deserialize message
	var m *discord.DiscordMessage
	err := json.Unmarshal(msg.Value, &m)
	if err != nil {
		slog.Error("error unmarshalling message %s", err)
		span.RecordError(err)
		return err
	}

	// call discord
	slog.Info("sending message to discord %v", m)
	err = p.discordHandler.SendMessage(ctx, m)
	if err != nil {
		slog.Error("error sending message to discord %s", err)
		span.RecordError(err)
		return err
	}

	slog.Info("send a message to discord %s", msg.Key)
	return nil
}

func (p *DiscordMessageProcessor) WriteTestMessage() error {
	ctx := context.Background()
	msg := &discord.DiscordMessage{
		Message: "test",
		Channel: "test",
	}

	d, err := json.Marshal(msg)
	if err != nil {
		slog.Error("error marshalling message %s", err)
		return err
	}

	km := kafkago.Message{
		Key:   []byte("test"),
		Value: d,
	}

	return p.kafkaProcessor.WriteMessage(ctx, km)
}
