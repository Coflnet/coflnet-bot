package db

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type Message struct {
	gorm.Model
	UserID    uint `gorm:"index"`
	Source    MessageSource
	Content   string
	Timestamp time.Time
}

type MessageSource string

const (
	MessageSourceDiscord MessageSource = "discord"
	MessageSourceRedis   MessageSource = "redis"
)

func SaveMessage(ctx context.Context, message *Message) error {
	ctx, span := tracer.Start(ctx, "save-message")
	defer span.End()

	slog.Debug("Saving message", "message", message)
	result := db.Create(message)

	if result.Error != nil {
		return result.Error
	}

	slog.Debug("Message saved")
	return nil
}
