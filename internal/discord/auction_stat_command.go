package discord

import (
	"encoding/json"
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/metrics"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func auctionStatCommand() *discordgo.ApplicationCommand {
	var zero float64 = 0
	var chatPerms int64 = discordgo.PermissionSendMessages

	return &discordgo.ApplicationCommand{
		Name:                     "auction-count",
		Description:              "Get the number of auctions in the last x minutes",
		DefaultMemberPermissions: &chatPerms,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "minutes",
				Description: "The number of minutes to look back",
				MinValue:    &zero,
				MaxValue:    60,
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "hours",
				Description: "The number of hours to look back",
				MinValue:    &zero,
				MaxValue:    24,
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "days",
				Description: "The number of days to look back",
				MinValue:    &zero,
				MaxValue:    30,
				Required:    false,
			},
		},
	}
}

func auctionStatCommmandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	minutes := 0
	if v, ok := optionMap["minutes"]; ok {
		fMinutes := v.Value.(float64)
		minutes = int(fMinutes)
	}

	hours := 0
	if v, ok := optionMap["hours"]; ok {
		fHours := v.Value.(float64)
		hours = int(fHours)
	}

	days := 0
	if v, ok := optionMap["days"]; ok {
		fDays := v.Value.(float64)
		days = int(fDays)
	}

	now := time.Now()
	start := now.Add(-time.Duration(minutes) * time.Minute).Add(-time.Duration(hours) * time.Hour).Add(-time.Duration(days) * time.Hour * 24)

	count, err := auctionsCount(start, now)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting auction count")
		metrics.ErrorOccurred()
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:         "Error getting auction count",
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
		return
	}

	message := fmt.Sprintf("%d auctions in the last %d minutes", count.Count, minutes)
	if hours > 0 {
		message = fmt.Sprintf("%d auctions in the last %d hours", count.Count, hours)
	}
	if days > 0 {
		message = fmt.Sprintf("%d auctions in the last %d days", count.Count, days)
	}

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         message,
			AllowedMentions: &discordgo.MessageAllowedMentions{},
		},
	})
}

func auctionsCount(start, end time.Time) (*AuctionCountResponse, error) {
	d := end.Sub(start)

	minutes := int(d.Minutes())

	url := fmt.Sprintf("%s/api/new-auctions?duration=%d", os.Getenv("AUCTION_STATS_URL"), minutes)

	resp, err := http.DefaultClient.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is not 200")
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data *AuctionCountResponse
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

type AuctionCountResponse struct {
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
	Count int64     `json:"count"`
}
