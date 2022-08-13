package coflnet

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	url2 "net/url"
	"os"
)

func PlayerName(uuid string) (string, error) {

	url, err := url2.JoinPath(PlayerNameHost(), "/PlayerName/name/", uuid)
	if err != nil {
		log.Error().Err(err).Msgf("can not create url for playername request")
		return "", err
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msgf("can not get playername from %s", url)
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msgf("can not close body")
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		log.Error().Msgf("can not get playername from %s, status code %d", url, resp.StatusCode)
		return "", fmt.Errorf("can not get playername from %s, status code %d", url, resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msgf("can not read body")
		return "", err
	}

	return string(b), nil
}

func PlayerNameHost() string {
	h := os.Getenv("PLAYER_NAME_HOST")

	if h == "" {
		log.Panic().Msgf("PLAYER_NAME_HOST not set")
	}

	return h
}
