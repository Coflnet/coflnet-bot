package coflnet

import (
	"time"

	"github.com/rs/zerolog/log"
)

func StartRefresh() {
	go refreshUsers()
}

func refreshUsers() {
	id := 1
	errorCounter := 0

	// refresh users, when 100 errors in a row happen return
	for errorCounter < 100 {

		err := refreshUser(id)
		if err != nil {
			errorCounter++
		} else {
			errorCounter = 0
		}

		if id%100 == 0 {
			log.Info().Msgf("refreshed %d users", id)
		}

		id++
		time.Sleep(time.Second * 20)
	}
}

func refreshUser(id int) error {
	_, err := LoadUserById(id)
	return err
}
