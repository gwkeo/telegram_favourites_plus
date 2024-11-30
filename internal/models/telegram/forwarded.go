package telegram

type Forwarded struct {
	Ok     bool       `json:"ok"`
	Result ForumTopic `json:"result"`
}

type ForumTopic struct {
	MessageThreadID int    `json:"message_thread_id"`
	Name            string `json:"name"`
}
