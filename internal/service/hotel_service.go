package service

import (
	"booking/internal/model"
	"booking/internal/repository"
	"errors"

	"github.com/google/uuid"
)

type HotelService struct {
	repo repository.HotelRepository
}

func InitHotelService(repo repository.HotelRepository) *HotelService {
	return &HotelService{repo: repo}
}

func (s *HotelService) GetByID(hotelUUID uuid.UUID) (*model.Hotel, error) {
	if hotelUUID == uuid.Nil {
		return nil, errors.New("invalid UUID")
	}
	hotel, err := s.repo.GetByID(hotelUUID)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}

func (s *HotelService) Create(hotel *model.Hotel) error {
	if hotel == nil {
		return errors.New("hotel cannot be nil")
	}
	err := s.repo.Create(hotel)
	if err != nil {
		return err
	}
	return nil
}

func (s *HotelService) Delete(hotelUUID uuid.UUID) error {
	if hotelUUID == uuid.Nil {
		return errors.New("invalid UUID")
	}
	err := s.repo.Delete(hotelUUID)
	if err != nil {
		return err
	}
	return nil
}
