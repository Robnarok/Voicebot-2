package model

import (
	"github.com/google/uuid"
)

type Channel struct {
	ID           uint32
	TextName     string
	VoiceName    string
	CategoryName string
}

func CreateChannel(textName, voiceName, categoryName string) Channel {
	u := uuid.New()
	channel := Channel{
		ID:           u.ID(),
		TextName:     textName,
		VoiceName:    voiceName,
		CategoryName: categoryName,
	}
	return channel
}
