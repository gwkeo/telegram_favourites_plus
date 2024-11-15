package models

type Type int

const (
	TextType Type = iota
	AnimationType
	PhotoType
	DocumentType
	VideoType
	VoiceType
	VideoNoteType
)

type Request struct {
	Id            int
	ForwardToChat int
	FromChat      int
	Type          Type
}

func ForwardMessageRequest(msgType Type, fromChatId, forwardToChat, id int) *Request {
	return &Request{
		Id:            id,
		ForwardToChat: forwardToChat,
		FromChat:      fromChatId,
		Type:          msgType,
	}
}
