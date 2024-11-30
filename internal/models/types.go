package models

type Type int

const (
	TrashType Type = iota
	TextType
	AnimationType
	PhotoType
	DocumentType
	VideoType
	VoiceType
	VideoNoteType
)
