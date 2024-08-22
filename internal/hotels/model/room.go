package model

import "github.com/google/uuid"

type Room struct {
	RoomID uuid.UUID `json:"room_id"`
	Number int       `json:"number"`
}
