package models

type Response struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result,omitempty"`
}

type Result struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Id                 int                `json:"message_id"`
	Chat               Chat               `json:"chat"`
	NewChatParticipant NewChatParticipant `json:"new_chat_participant"`
	Text               string             `json:"text,omitempty"`
	Animation          Animation          `json:"animation,omitempty"`
	Photo              []Photo            `json:"photo,omitempty"`
	Document           Document           `json:"document,omitempty"`
	Video              Video              `json:"video,omitempty"`
	Voice              Voice              `json:"voice,omitempty"`
	VideoNote          VideoNote          `json:"video_note,omitempty"`
	Entities           Entities           `json:"entities,omitempty"`
}

type Chat struct {
	Id int `json:"id"`
}

type NewChatParticipant struct {
	Id       int    `json:"id"`
	IsBot    bool   `json:"is_bot"`
	Username string `json:"username"`
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
