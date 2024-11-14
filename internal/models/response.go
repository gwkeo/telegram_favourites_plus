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
	Id        int       `json:"message_id"`
	Chat      Chat      `json:"chat"`
	Text      string    `json:"text,omitempty"`
	Animation *struct{} `json:"animation,omitempty"`
	Photo     *struct{} `json:"photo,omitempty"`
	Document  *struct{} `json:"document,omitempty"`
	Video     *struct{} `json:"video,omitempty"`
	Voice     *struct{} `json:"voice,omitempty"`
	VideoNote *struct{} `json:"video_note,omitempty"`
	Caption   string    `json:"caption,omitempty"`
}

type Chat struct {
	Id int `json:"id"`
}
