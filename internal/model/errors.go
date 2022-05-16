package model

import "fmt"

type UserNotFoundError struct {
	UserId int
}

func (u *UserNotFoundError) Error() string {
	if u.UserId != 0 {
		return fmt.Sprintf("user %d not found", u.UserId)
	}
	return "user not found"
}
