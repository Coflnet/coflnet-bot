package discord

import (
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
)

func FlipTracked(trackedFlip *model.FlipTrack) error {

	_, err := session.ChannelMessageSendComplex(flipsChannelId(), &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: trackedFlip.Seller,
				Author: &discordgo.MessageEmbedAuthor{
					Name: trackedFlip.Seller,
					// IconURL: iconUrl,
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "price",
						Value: strconv.Itoa(int(trackedFlip.Sell.SellPrice)),
					},
					{
						Name:  "auctionUuid",
						Value: trackedFlip.OriginAuction,
					},
					{
						Name:  "time",
						Value: trackedFlip.FoundTime.Format(time.RFC3339),
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
