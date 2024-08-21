package model

import "github.com/google/uuid"

type POI struct {
	POIID       uuid.UUID `json:"poi_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
