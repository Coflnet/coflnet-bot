package utils

import (
  "os"
  "github.com/rs/zerolog/log"
)

func Env(key string) string {
  value, ok := os.LookupEnv(key)
  if !ok {
    log.Panic().Msgf("missing %s environment variable", key)
  }
  return value
}
