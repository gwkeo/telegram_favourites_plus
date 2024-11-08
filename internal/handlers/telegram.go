package handlers

import (
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
	"sync"
)

type Type int

const (
	Text Type = iota
	Animation
)

func HandleResults(results []models.Result) {
	wg := &sync.WaitGroup{}
	for _, r := range results {
		wg.Add(1)
		go HandlerResult(r, wg)
	}
}

func HandlerResult(r models.Result, wg *sync.WaitGroup) Type {
	defer wg.Done()
	if r.Message.Text != "" {
		return Text
	} else {
		return Animation
	}
}
