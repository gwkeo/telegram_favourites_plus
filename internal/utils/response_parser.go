package utils

import (
	"encoding/json"
	"errors"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
)

func Response(body []byte) (*models.Response, error) {
	response := &models.Response{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return response, errors.New("error parsing response: " + err.Error())
	}
	return response, nil
}
