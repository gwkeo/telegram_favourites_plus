package models

type GetUpdatesResponse struct {
	Status bool     `json:"ok"`
	Result []Result `json:"result"`
}

type Result struct {
	Message  Message `json:"message"`
	UpdateId int     `json:"update_id"`
}

type Message struct {
	Chat   Chat    `json:"chat"`
	Text   string  `json:"text,omitempty"`
	Photos []Photo `json:"photos,omitempty"`
}

type Chat struct {
	ChatId int `json:"message_id"`
}

type Photo struct {
	Id string `json:"file_id"`
}
