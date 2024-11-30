package utils

import (
	"encoding/json"
	"errors"
	"github.com/gwkeo/telegram_favourites_plus/internal/models/telegram"
)

func ParseUpdates(body []byte) (*telegram.NewUpdates, error) {
	response := &telegram.NewUpdates{}
	if err := json.Unmarshal(body, &response); err != nil {
		return response, errors.New("error parsing response: " + err.Error())
	}
	return response, nil
}

func ParseCreated(body []byte) (*telegram.Forwarded, error) {
	forwarded := &telegram.Forwarded{}
	if err := json.Unmarshal(body, &forwarded); err != nil {
		return nil, errors.New("error parsing forum topic: " + err.Error())
	}
	return forwarded, nil
}
