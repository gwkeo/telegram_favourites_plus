package models

type Branch struct {
	ID      int
	Type    Type
	ForumID int
}

var TopicNames = map[string]Type{
	"Мусорка":             TrashType,
	"Текст":               TextType,
	"Гифки":               AnimationType,
	"Фотографии":          PhotoType,
	"Документы":           DocumentType,
	"Видео":               VideoType,
	"Голосовые сообщения": VoiceType,
	"Кружочки":            VideoNoteType,
}
