package utils

import "github.com/Coflnet/coflnet-bot/internal/model"

func FilterUsersForPreferredUsers(userId string, users []*model.User) *model.User {

    if len(users) == 1 {
        if users[0].PreferredUsername != "" {
            return users[0]
        }
    }

	for _, user := range users {
		if user.PreferredDiscordId == userId {
			return user
		}
	}

	return users[0]
}
