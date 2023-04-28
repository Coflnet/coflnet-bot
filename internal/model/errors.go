package model

import "fmt"

type UserNotFoundError struct {
	UserId    int
	DiscordId string
	UUID      string
}

func (u *UserNotFoundError) Error() string {
	if u.UserId != 0 {
		return fmt.Sprintf("user with id %d or discord id %s or uuid %s not found", u.UserId, u.DiscordId, u.UUID)
	}
	return "user not found"
}
