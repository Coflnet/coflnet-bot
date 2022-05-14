package coflnet

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const PREMIUM_PRODUCT_SLUG = "premium"

func PaymentUserById(userId int) (time.Time, error) {
	url := fmt.Sprintf("%s/User/%d/owns/%s/until", os.Getenv("PAYMENT_URL"), userId, PREMIUM_PRODUCT_SLUG)

	response, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Error().Err(err).Msgf("error getting user from payment, userId: %d", userId)
		return time.Now(), err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return time.Now(), err
	}

	t := strings.ReplaceAll(string(body), "\"", "")

	const layout = "2006-01-02T15:04:05"
	result, err := time.Parse(layout, t)

	if err != nil {
		log.Error().Err(err).Msgf("could not parse premium time for user %d", userId)
		return time.Now(), err
	}

	return result, nil
}
