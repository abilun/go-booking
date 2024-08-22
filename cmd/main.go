package main

import (
	"booking/internal/repository/cassandra"
	"booking/internal/service"
	"booking/internal/transport/rest"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
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

	hotelRepo := cassandra.InitHotelRepository(session)
	hotelService := service.InitHotelService(hotelRepo)

	hotelHandler := rest.NewHotelHandler(hotelService)
	mappers := []rest.Mapper{
		{
			Method:  "GET",
			Path:    "/hotel/{uuid}",
			Handler: hotelHandler.GetHotel,
		},
		{
			Method:  "DELETE",
			Path:    "/hotel/{uuid}",
			Handler: hotelHandler.DeleteHotel,
		},
		{
			Method:  "POST",
			Path:    "/hotel",
			Handler: hotelHandler.CreateHotel,
		},
	}
	handlerRouter := rest.NewRouter(mappers)

	err = http.ListenAndServe("localhost:8080", &handlerRouter.Mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}
