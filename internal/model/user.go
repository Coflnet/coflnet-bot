package model

import "time"

type OwnedProducts struct {
	ProductSlug string    `json:"slug" bson:"slug"`
	ExpiresAt   time.Time `json:"expiresAt" bson:"expires_at"`

	// maybe rename to Role
	Color string `json:"color" bson:"color"`
}

type User struct {

	// the id given by cofl
	UserId int `json:"userId" bson:"user_id"`

	PremiumUntil time.Time `json:"premiumUntil" bson:"premium_until"`

	// a deprecated list of discord usernames + discriminator
	DiscordNames []string `json:"discordNames" bson:"discord_names"`

	// a list of minecraft uuids associated with this user
	MinecraftUuids []string `json:"minecraftUuids" bson:"minecraft_uuids"`

	// the last time the user was refreshed
	LastRefresh time.Time `json:"lastRefresh" bson:"last_refresh"`

	// a checkmark if the user has the flipper role
	HasFlipperRole bool `json:"hasFlipperRole" bson:"has_flipper_role"`

	// user can set a preferred username of all his discord accounts
	PreferredUsername string `json:"preferredUsername" bson:"preferred_username"`

	// user can set a preferred uuid of all his minecraft accounts
	PreferredUUID string `json:"preferredUUID" bson:"preferred_uuid"`

	// a new list with the actual discord ids, should replace the DiscordNames
	DiscordIds []string `json:"discordIds" bson:"discord_ids"`

	// the preferred discord id
	PreferredDiscordId string `json:"preferredDiscordId" bson:"preferred_discord_id"`

	// a list of all products the user owns
	OwnedProducts []OwnedProducts `json:"ownedProducts" bson:"owned_products"`
}

func (u *User) HasPremium() bool {
	return u.PremiumUntil.After(time.Now())
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
