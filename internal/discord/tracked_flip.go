package discord

import (
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func FlipTracked(flip *model.Flip) error {

	_, err := session.ChannelMessageSendComplex(flipsChannelId(), &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: flip.Buy.ItemName,
				Author: &discordgo.MessageEmbedAuthor{
					Name: flip.Buy.ProfileID,
					// IconURL: flip.Buy.,
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "price",
						Value: strconv.Itoa(flip.Sell.HighestBidAmount),
					},
					{
						Name:  "profit",
						Value: strconv.Itoa(flip.Profit),
					},
				},
				Type: "rich",
			},
		},
	})

	if err != nil {
		log.Error().Err(err).Msgf("error sending message to discord chat")
		return err
	}

	return nil
}

func flipsChannelId() string {
	flipsChannel := os.Getenv("FLIPS_CHANNEL")

	if flipsChannel == "" {
		log.Panic().Msg("FLIPS_CHANNEL id not set")
	}
	return flipsChannel
}
