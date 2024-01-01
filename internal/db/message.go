package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type Message struct {
	gorm.Model
	Content         string
	Timestamp       time.Time
	DiscordUserId   string
	DiscordUsername string
	IsBot           bool
	ChannelId       string
	GuildId         string
	ThreadId        *string
}

func SaveMessage(ctx context.Context, message *Message) error {
	ctx, span := tracer.Start(ctx, "save-message")
	defer span.End()

	result := db.Save(message)
	if result.Error != nil {
		span.RecordError(result.Error)
		return result.Error
	}

	slog.Info(fmt.Sprintf("saved message from %s", message.DiscordUsername))
	return nil
}
