package model

import "time"

type OwnedProducts struct {
    ProductSlug string `json:"slug" bson:"slug"`
    ExpiresAt time.Time `json:"expiresAt" bson:"expires_at"`

    // maybe rename to Role
    Color string `json:"color" bson:"color"`
}

type User struct {
	UserId         int       `json:"userId" bson:"user_id"`
	PremiumUntil   time.Time `json:"premiumUntil" bson:"premium_until"`
	DiscordNames   []string  `json:"discordNames" bson:"discord_names"`
	MinecraftUuids []string  `json:"minecraftUuids" bson:"minecraft_uuids"`

	LastRefresh    time.Time `json:"lastRefresh" bson:"last_refresh"`
	HasFlipperRole bool      `json:"hasFlipperRole" bson:"has_flipper_role"`

	PreferredUsername string `json:"preferredUsername" bson:"preferred_username"`
	PreferredUUID     string `json:"preferredUUID" bson:"preferred_uuid"`

	DiscordIds         []string `json:"discordIds" bson:"discord_ids"`
	PreferredDiscordId string   `json:"preferredDiscordId" bson:"preferred_discord_id"`

    OwnedProducts []OwnedProducts `json:"ownedProducts" bson:"owned_products"`
}

func (u *User) HasPremium() bool {
	return u.PremiumUntil.After(time.Now())
}

func (u *User) Username() string {
	if u.PreferredUsername != "" {
		return u.PreferredUsername
	}

	if len(u.DiscordNames) > 0 {
		return u.DiscordNames[0]
	}

	return ""
}

func (u *User) UUID() string {
	if u.PreferredUUID != "" {
		return u.PreferredUUID
	}

	if len(u.MinecraftUuids) > 0 {
		return u.MinecraftUuids[0]
	}

	return ""
}
