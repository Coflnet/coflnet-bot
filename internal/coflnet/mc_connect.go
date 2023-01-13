package coflnet

import (
	"encoding/json"
	"fmt"
	"github.com/Coflnet/coflnet-bot/internal/model"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

type McConnectGetUserResponse struct {
	ExternalID string `json:"externalId"`
	Accounts   []struct {
		AccountUUID string `json:"accountUuid"`
		Verified    bool   `json:"verified"`
		UpdatedAt   string `json:"updatedAt"`
	} `json:"accounts"`
}

func UserMcConnect(userId int) (*model.User, error) {
	url := fmt.Sprintf("%s/Connect/user/%d", os.Getenv("MC_CONNECT_URL"), userId)

	response, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Error().Err(err).Msgf("error getting user from mc connect, userId: %d", userId)
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msgf("error reading response body")
		return nil, err
	}

	var result McConnectGetUserResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error().Err(err).Msgf("error unmarshalling response body")
		return nil, err
	}

	var uuids []string
	for _, u := range result.Accounts {

		// only add the verified ones
		if !u.Verified {
			continue
		}

		uuids = append(uuids, u.AccountUUID)
	}

	user := model.User{
		UserId:         userId,
		MinecraftUuids: uuids,
	}

	return &user, nil
}

func UserMcConnectByUUID(uuid string) (*model.User, error) {
  log.Info().Msgf("checking mc connect user by the uuid: %s", uuid)
	url := fmt.Sprintf("%s/Connect/minecraft/%s", os.Getenv("MC_CONNECT_URL"), uuid)

	response, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Error().Err(err).Msgf("error getting user from mc connect, uuid: %s", uuid)
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msgf("error reading response body")
		return nil, err
	}

	var result McConnectGetUserResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error().Err(err).Msgf("error unmarshalling response body")
		return nil, err
	}

	var uuids []string
	for _, u := range result.Accounts {

		// only add the verified ones
		if !u.Verified {
			continue
		}

		uuids = append(uuids, u.AccountUUID)
	}

  id, err := strconv.Atoi(result.ExternalID)
  if err != nil {
    log.Error().Err(err).Msgf("error converting external id to int")
    return nil, err
  }

	user := model.User{
		UserId:         id,
		MinecraftUuids: uuids,
	}

	return &user, nil
}

func GetUsersFromId(amount, offset int) ([]*model.User, error) {
	url := fmt.Sprintf("%s/Connect/users?amount=%d&offset=%d", os.Getenv("MC_CONNECT_URL"), amount, offset)

	response, err := http.DefaultClient.Get(url)

	if err != nil {
		log.Error().Err(err).Msgf("there was an error when getting users from mc connect, url: %s", url)
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msgf("error reading response body")
		return nil, err
	}

	var result []*McConnectGetUserResponse
	err = json.Unmarshal(body, &result)

	var users []*model.User

	for _, u := range result {

		var uuids []string
		for _, u := range u.Accounts {
			uuids = append(uuids, u.AccountUUID)
		}

		i, err := strconv.Atoi(u.ExternalID)
		if err != nil {
			log.Error().Err(err).Msgf("error converting external id to int")
			continue
		}

		user := model.User{
			UserId:         i,
			MinecraftUuids: uuids,
		}

		users = append(users, &user)
	}

	return users, nil
}
