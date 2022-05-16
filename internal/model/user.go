package model

import "time"

type User struct {
	UserId         int       `json:"userId" bson:"user_id"`
	PremiumUntil   time.Time `json:"premiumUntil" bson:"premium_until"`
	DiscordNames   []string  `json:"discordNames" bson:"discord_names"`
	MinecraftUuids []string  `json:"minecraftUuids" bson:"minecraft_uuids"`
	LastRefresh    time.Time `json:"lastRefresh" bson:"last_refresh"`
	HasFlipperRole bool      `json:"hasFlipperRole" bson:"has_flipper_role"`
}

func (u *User) HasPremium() bool {
	return u.PremiumUntil.After(time.Now())
}
