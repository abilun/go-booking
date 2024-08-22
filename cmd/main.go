package main

import (
	"booking/internal/hotels/repository/cassandra"
	"booking/internal/hotels/services"
	rest2 "booking/internal/hotels/webapi"
	"booking/internal/infra/webapi"
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
	hotelService := services.InitHotelService(hotelRepo)

	hotelHandler := rest2.NewHotelHandler(hotelService)
	mappers := []webapi.Mapper{
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
	handlerRouter := webapi.NewRouter(mappers)

	err = http.ListenAndServe("localhost:8080", &handlerRouter.Mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}
