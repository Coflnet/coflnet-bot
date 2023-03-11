package coflnet

import (
	"errors"

	"github.com/Coflnet/coflnet-bot/internal/model"
)

// UserById loads a user, potentially cached
func LoadUserById(id int) (*model.User, error) {
    return nil, errors.New("not implemented")
}

func LoadUserByUUID(uuid string) (*model.User, error) {
    return nil, errors.New("not implemented")
}

func LoadUser(mcUser *model.User) (*model.User, error) {
    return nil, errors.New("not implemented")
}
