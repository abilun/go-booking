package model

import (
	"github.com/google/uuid"
)

type Hotel struct {
	HotelID     uuid.UUID `json:"hotel_id"`
	Name        string    `json:"name"`
	Address     Address   `json:"address"`
	Description string    `json:"description"`
	Phone       string    `json:"phone"`
	POIs        []POI     `json:"pois"`
	Rooms       []Room    `json:"rooms"`
}
