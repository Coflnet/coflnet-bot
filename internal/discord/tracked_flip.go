package discord

import (
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/coflnet"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/rs/zerolog/log"
	"os"
)

func FlipTracked(flip *model.Flip) error {

	playerName, err := coflnet.PlayerName(flip.Buy.AuctioneerID)
	if err != nil {
		log.Error().Err(err).Msgf("could not get player name for %s", flip.Buy.AuctioneerID)
		return err
	}

	_, err = session.ChannelMessageSendComplex(flipsChannelId(), &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: flip.Buy.ItemName,
				Author: &discordgo.MessageEmbedAuthor{
					Name:    playerName,
					IconURL: iconUrl(flip),
				},
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "price",
						Value: humanize.Comma(int64(flip.Sell.HighestBidAmount)),
					},
					{
						Name:  "profit",
						Value: humanize.Comma(int64(flip.Profit)),
					},
					{
						Name:  "Time",
						Value: fmt.Sprintf("<t:%d:f>", flip.Sell.End.Unix()),
					},
					{
						Name:  "Finder",
						Value: finder(flip.Finder),
					},
					{
						Name:  "Sell UUID",
						Value: flip.Sell.UUID,
					},
					{
						Name:  "Buy UUID",
						Value: flip.Buy.UUID,
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

func finder(finder interface{}) string {
	if finder == nil {
		return "-"
	}

	return fmt.Sprintf("%s", finder)
}

func iconUrl(flip *model.Flip) string {
	return fmt.Sprintf("https://crafatar.com/avatars/%s", flip.Sell.AuctioneerID)
}

func flipsChannelId() string {
	flipsChannel := os.Getenv("FLIPS_CHANNEL")

	if flipsChannel == "" {
		log.Panic().Msg("FLIPS_CHANNEL id not set")
	}
	return flipsChannel
}
