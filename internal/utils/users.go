package utils

import "github.com/Coflnet/coflnet-bot/internal/model"

func FilterUsersForPreferredUsers(userId string, users []*model.User) []*model.User {
    var filteredUsers []*model.User
    for _, user := range users {
        if user.PreferredDiscordId  == userId {
            filteredUsers = append(filteredUsers, user)
        }
    }

    if len(filteredUsers) == 0 {
        return users
    }

    return filteredUsers
}
