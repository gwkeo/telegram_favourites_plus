package utils

import (
	"errors"
	"flag"
)

func Api() (string, error) {
	apiKey := flag.String("api-key", "", "Telegram API Key")
	flag.Parse()
	if *apiKey == "" {
		return "", errors.New("apikey is required")
	}
	return *apiKey, nil
}
