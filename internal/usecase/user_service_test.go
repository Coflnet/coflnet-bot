package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleConvertRedisMessage(t *testing.T) {
	service, err := NewUserService("", "", "")
	if err != nil {
		panic(err)
	}

	redisMessageContent := "{\"externalId\":\"7\",\"accounts\":[]}"
	ctx := context.Background()

	externalId, err := service.LoadExternalId(ctx, redisMessageContent)

	if err != nil {
		panic(err)
	}

	// assert that the externalId has the correct external id
	assert.Equal(t, "7", externalId)
}
