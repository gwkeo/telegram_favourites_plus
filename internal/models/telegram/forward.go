package telegram

import "github.com/gwkeo/telegram_favourites_plus/internal/models"

type Forward struct {
	ID       int
	FromChat int
	ThreadId int
	Type     models.Type
}
