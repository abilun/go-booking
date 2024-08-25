package webapi

import (
	"booking/internal/hotels/model"
	"booking/internal/hotels/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type HotelHandler struct {
	HotelService services.HotelService
}

func NewHotelHandler(service *services.HotelService) *HotelHandler {
	return &HotelHandler{
		HotelService: *service,
	}
}

func (h *HotelHandler) GetHotel(w http.ResponseWriter, req *http.Request) {

	hotelUUID, err := uuid.Parse(req.PathValue("uuid"))
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	hotel, err := h.HotelService.GetByID(hotelUUID)
	if err != nil {
		http.Error(w, "Hotel not found", http.StatusNotFound)
		return
	}

	json, err := json.Marshal(hotel)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func (h *HotelHandler) DeleteHotel(w http.ResponseWriter, req *http.Request) {
	hotelUUID, err := uuid.Parse(req.PathValue("uuid"))
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	err = h.HotelService.Delete(hotelUUID)
	if err != nil {
		http.Error(w, "Hotel not found", http.StatusNotFound)
		return
	}
}

func (h *HotelHandler) CreateHotel(w http.ResponseWriter, req *http.Request) {
	var hotel model.Hotel
	err := json.NewDecoder(req.Body).Decode(&hotel)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = h.HotelService.Create(&hotel)
	if err != nil {
		log.Println("Error creating hotel:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
