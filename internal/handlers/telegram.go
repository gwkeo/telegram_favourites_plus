package handlers

import (
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
	"sync"
)

type Type int

const (
	Text Type = iota
	Animation
	Photo
	Document
	Video
	Voice
	VideoNote
)

func TypeOfResult(r models.Result, wg *sync.WaitGroup) Type {
	defer wg.Done()
	if r.Message.Text != "" {
		return Text
	} else if r.Message.Animation != nil {
		return Animation
	} else if r.Message.Photo != nil {
		return Photo
	} else if r.Message.Document != nil {
		return Document
	} else if r.Message.Voice != nil {
		return Voice
	} else if r.Message.Video != nil {
		return Video
	} else {
		return VideoNote
	}
}

func HandleTelegramResponse(results []models.Result) {
	wg := &sync.WaitGroup{}
	for _, r := range results {
		wg.Add(1)
		go func() {
			resType := TypeOfResult(r, wg)
		}()
	}

	wg.Wait()
}
