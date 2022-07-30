package model

import (
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Warning struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	User   *discordgo.User    `json:"user" bson:"user"`
	Warner *discordgo.Member  `json:"warner" bson:"warner"`
	Until  time.Time          `json:"duration" bson:"duration"`
	Reason string             `json:"reason" bson:"reason"`
}
