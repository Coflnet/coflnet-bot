package utils

import "github.com/Coflnet/coflnet-bot/internal/model"

func FilterUsersForPreferredUsers(userId string, users []*model.User) *model.User {
	for _, user := range users {
		if user.PreferredDiscordId == userId {
			return user
		}
	}

	return users[0]
}
