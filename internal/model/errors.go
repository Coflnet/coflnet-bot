package model

import "fmt"

type UserNotFoundError struct {
	UserId    int
	DiscordId string
}

func (u *UserNotFoundError) Error() string {
	if u.UserId != 0 {
		return fmt.Sprintf("user with %d or discord id %s not found", u.UserId, u.DiscordId)
	}
	return "user not found"
}
