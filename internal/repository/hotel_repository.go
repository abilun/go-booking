package repository

import (
	"booking/internal/model"

	"github.com/google/uuid"
)

type HotelRepository interface {
	GetByID(hotelUUID uuid.UUID) (*model.Hotel, error)
	Create(user *model.Hotel) error
	Delete(hotelUUID uuid.UUID) error
}
