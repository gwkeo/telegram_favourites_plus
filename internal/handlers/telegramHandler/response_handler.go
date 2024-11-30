package telegramHandler

import (
	"github.com/gwkeo/telegram_favourites_plus/internal/events"
	"github.com/gwkeo/telegram_favourites_plus/internal/models"
	"github.com/gwkeo/telegram_favourites_plus/internal/models/telegram"
)

func MsgType(r *telegram.Message) models.Type {
	if r.Text != "" {
		return models.TextType
	} else if r.Animation != nil {
		return models.AnimationType
	} else if r.Photo != nil {
		return models.PhotoType
	} else if r.Document != nil {
		return models.DocumentType
	} else if r.Voice != nil {
		return models.VoiceType
	} else if r.Video != nil {
		return models.VideoType
	} else {
		return models.VideoNoteType
	}
}

func EventType(upd telegram.Result) events.Types {
	if upd.MyChatMember != nil && upd.MyChatMember.NewChatMember.CanManageTopics {
		return events.BotAdded
	} else if upd.Message != nil {
		return events.MessageType
	}
	return events.Default
}
