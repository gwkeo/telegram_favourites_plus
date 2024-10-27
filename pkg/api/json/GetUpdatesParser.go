package json

import (
	"encoding/json"
	"github.com/gwkeo/telegram_favourites_plus/pkg/api/json/models"
)

func ParseGetUpdatesResponse(response []byte) (*models.GetUpdatesResponse, error) {
	jsonResponse := &models.GetUpdatesResponse{}
	ok := json.Unmarshal(response, jsonResponse)
	if ok != nil {
		return &models.GetUpdatesResponse{}, ok
	}
	return jsonResponse, nil
}
