package telegram

type NewUpdates struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result,omitempty"`
}

type Result struct {
	UpdateID     int           `json:"update_id"`
	Message      *Message      `json:"message"`
	MyChatMember *MyChatMember `json:"my_chat_member"`
}

type Message struct {
	ID        int        `json:"message_id"`
	Chat      Chat       `json:"chat"`
	Text      string     `json:"text,omitempty"`
	Animation *Animation `json:"animation,omitempty"`
	Photo     *[]Photo   `json:"photo,omitempty"`
	Document  *Document  `json:"document,omitempty"`
	Video     *Video     `json:"video,omitempty"`
	Voice     *Voice     `json:"voice,omitempty"`
	VideoNote *VideoNote `json:"video_note,omitempty"`
	Entities  *Entities  `json:"entities,omitempty"`
}

type MyChatMember struct {
	NewChatMember *NewChatMember `json:"new_chat_member,omitempty"`
	Chat          Chat           `json:"chat"`
}

type Chat struct {
	ID int `json:"id"`
}

type NewChatMember struct {
	CanManageTopics bool `json:"can_manage_topics"`
}

type Entities struct {
	Type string `json:"type"`
}

type Animation struct{}
type Photo struct{}
type Document struct{}
type Video struct{}
type Voice struct{}
type VideoNote struct{}
