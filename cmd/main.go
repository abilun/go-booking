package main

import (
	"booking/internal/model"
	"booking/internal/repository/cassandra"
	"booking/internal/service"
	"log"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func main() {
	hotelCluster := gocql.NewCluster("localhost")
	hotelCluster.Keyspace = "hotel"
	hotelCluster.Consistency = gocql.Quorum
	session, err := hotelCluster.CreateSession()

	if err != nil {
		panic(err)
	}
	defer session.Close()

	hotelUuid := uuid.New()

	hotelRepo := cassandra.InitHotelRepository(session)
	hotelService := service.InitHotelService(hotelRepo)

	hotelService.Create(&model.Hotel{
		HotelID: hotelUuid,
		Name:    "Hotel Name",
		Address: model.Address{
			Country:  "Poland",
			City:     "Warsaw",
			Street:   "Marszałkowska",
			Building: 1,
			Entrance: 1,
			ZipCode:  "00-000",
		},
		Description: "The best hotel in the world",
		Phone:       "123456789",
	})

	hotel, err := hotelRepo.GetByID(hotelUuid)
	if err != nil {
		panic(err)
	}

	hotelService.Delete(hotelUuid)

	log.Printf("hotel: %v\n", hotel)
}
