package db

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log/slog"
	"time"
)

type User struct {
	gorm.Model

	// refers to the main user id, usually the external id
	ExternalId string `gorm:"uniqueIndex"`

	DiscordAccounts   []DiscordAccount
	MinecraftAccounts []MinecraftAccount

	// if the user has premium this timestamp is in the future
	PremiumUntil *time.Time

	// if the user has premium plus this timestamp is in the future
	PremiumPlusUntil *time.Time

	// timestamp when the information of the user was last updated
	LastUpdated *time.Time
}

type DiscordAccount struct {
	gorm.Model
	DiscordID string `gorm:"uniqueIndex"`
	UserID    uint   `gorm:"index"`
	Preferred bool
}

type MinecraftAccount struct {
	gorm.Model
	MinecraftUUID string `gorm:"uniqueIndex"`
	Preferred     bool
	UserID        uint `gorm:"index"`
}

type UserHasNoMinecraftAccountError struct {
	UserId     uint
	ExternalId string
}

func (e *UserHasNoMinecraftAccountError) Error() string {
	return fmt.Sprintf("user with id %d and external id %s has no minecraft account", e.UserId, e.ExternalId)
}

type UserHasNoDiscordAccountError struct {
	UserId     uint
	ExternalId string
}

func (e *UserHasNoDiscordAccountError) Error() string {
	return fmt.Sprintf("user with id %d and external id %s has no discord account", e.UserId, e.ExternalId)
}

func (u *User) HasPremium() bool {
	return u.PremiumUntil.After(time.Now())
}

func (u *User) HasPremiumPlus() bool {
	return u.PremiumPlusUntil.After(time.Now())
}

func (u *User) PreferredUUID() (string, error) {
	if len(u.MinecraftAccounts) == 0 {
		return "", errors.New("no minecraft accounts")
	}

	for _, acc := range u.MinecraftAccounts {
		if acc.Preferred {
			return acc.MinecraftUUID, nil
		}
	}

	return u.MinecraftAccounts[0].MinecraftUUID, nil
}

func (u *User) PreferredDiscordID() (string, error) {
	if len(u.DiscordAccounts) == 0 {
		return "", errors.New("no discord accounts")
	}

	for _, acc := range u.DiscordAccounts {
		if acc.Preferred {
			return acc.DiscordID, nil
		}
	}

	return u.DiscordAccounts[0].DiscordID, nil
}

func UserByExternalId(ctx context.Context, externalId string) (*User, error) {
	ctx, span := tracer.Start(ctx, "user-by-external-id")
	defer span.End()
	span.SetAttributes(attribute.String("externalId", externalId))

	user := &User{}
	result := db.
		Preload(clause.Associations).
		Where(&User{ExternalId: externalId}).
		First(user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			slog.Debug(fmt.Sprintf("user with external id %s not found", externalId))
			span.SetAttributes(attribute.Bool("found", false))
			return nil, nil
		}

		span.RecordError(result.Error)
		slog.Error("Error loading user", "err", result.Error)
		return nil, result.Error
	}

	slog.Debug(fmt.Sprintf("loaded user with external id %s", externalId))
	span.SetAttributes(attribute.Bool("found", true))
	return user, nil
}

func SaveUser(ctx context.Context, user *User) error {
	ctx, span := tracer.Start(ctx, "save-user")
	defer span.End()
	span.SetAttributes(attribute.String("externalId", user.ExternalId))

	result := db.Session(&gorm.Session{FullSaveAssociations: true}).Save(user)
	if result.Error != nil {
		return result.Error
	}

	slog.Debug(fmt.Sprintf("saved user with external id %s", user.ExternalId))
	return nil
}

func UserByDiscordId(ctx context.Context, id string) (*User, error) {
	ctx, span := tracer.Start(ctx, "user-by-discord-id")
	defer span.End()
	span.SetAttributes(attribute.String("discordId", id))

	user := &User{}
	result := db.
		Preload(clause.Associations).
		Joins("JOIN discord_accounts ON discord_accounts.user_id = users.id").
		Where("discord_accounts.discord_id = ?", id).
		First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
