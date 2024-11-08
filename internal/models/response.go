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
	Id        int      `json:"message_id"`
	Chat      Chat     `json:"chat"`
	Text      string   `json:"text,omitempty"`
	Animation struct{} `json:"animation,omitempty"`
}

type Chat struct {
	Id int `json:"id"`
}

type Animation struct {
}
