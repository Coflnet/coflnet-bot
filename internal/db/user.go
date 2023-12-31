package db

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type User struct {
	gorm.Model

	// refers to the main user id, usually the external id
	ExternalId string `gorm:"uniqueIndex"`

	DiscordAccounts []DiscordAccount
	HypixelAccounts []MinecraftAccount

	// a user can have multiple minecraft accounts
	// the preferred one is used for interactions
	PreferredMinecraftAccountID uint
	PreferredMinecraftAccount   *MinecraftAccount

	// a user can have multiple discord accounts
	// the preferred one is used for interactions
	PreferredDiscordAccountID uint
	PreferredDiscordAccount   *DiscordAccount

	// if the user has premium this timestamp is in the future
	PremiumUntil time.Time

	// if the user has premium plus this timestamp is in the future
	PremiumPlusUntil time.Time

	// timestamp when the information of the user was last updated
	LastUpdated time.Time
}

type DiscordAccount struct {
	gorm.Model
	DiscordID string

	UserID uint `gorm:"uniqueIndex"`
}

type MinecraftAccount struct {
	gorm.Model
	MinecraftUUID string `gorm:"uniqueIndex"`

	UserID uint
}

func (u *User) HasPremium() bool {
	return u.PremiumUntil.After(time.Now())
}

func (u *User) HasPremiumPlus() bool {
	return u.PremiumPlusUntil.After(time.Now())
}

func (u *User) PreferredUUID() (string, error) {
	if u.PreferredMinecraftAccount == nil {
		return "", errors.New("no preferred minecraft account")
	}

	return u.PreferredMinecraftAccount.MinecraftUUID, nil
}

func UserByExternalId(ctx context.Context, externalId string) (*User, error) {
	ctx, span := tracer.Start(ctx, "user-by-external-id")
	defer span.End()
	span.SetAttributes(attribute.String("externalId", externalId))

	user := &User{}
	result := db.Where(&User{ExternalId: externalId}).First(user)

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
	err := db.Model(&User{}).Where(DiscordAccount{DiscordID: id}).Association("DiscordAccounts").Find(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
