package handlers

import (
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
)

func TypeOfResult(r models.Message) models.Type {
	if r.Text != "" {
		return models.TextType
	} else if r.Animation != struct{}{} {
		return models.AnimationType
	} else if r.Photo != nil {
		return models.PhotoType
	} else if r.Document != struct{}{} {
		return models.DocumentType
	} else if r.Voice != struct{}{} {
		return models.VoiceType
	} else if r.Video != struct{}{} {
		return models.VideoType
	} else {
		return models.VideoNoteType
	}
}

func HandleTelegramResponse(results []models.Result, requests chan<- *models.Request) {
	for _, r := range results {
		resType := TypeOfResult(r.Message)
		req := models.ForwardMessageRequest(resType, r.Message.Chat.Id, r.Message.Chat.Id, r.Message.Id)
		requests <- req
	}
}
