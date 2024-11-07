package models

type Response struct {
	Result []Result `json:"result"`
}

type Result struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Id        int      `json:"message_id"`
	Chat      Chat     `json:"from"`
	Text      string   `json:"text"`
	Animation struct{} `json:"animation"`
}

type Chat struct {
	Id int `json:"id"`
}
