package cassandra

import (
	model2 "booking/internal/hotels/model"
	cassandraInfra "booking/internal/infra/cassandra"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type CassandraAddress struct {
	Country  string `cql:"country"`
	City     string `cql:"city"`
	Street   string `cql:"street"`
	Building int    `cql:"building"`
	Entrance int    `cql:"entrance"`
	ZipCode  string `cql:"zip_code"`
}

type CassandraHotel struct {
	HotelID     gocql.UUID       `cql:"hotel_id"`
	Name        string           `cql:"name"`
	Address     CassandraAddress `cql:"address"`
	Description string           `cql:"description"`
	Phone       string           `cql:"phone"`
	POIs        []gocql.UUID     `cql:"pois"`
	Rooms       []gocql.UUID     `cql:"rooms"`
}

// Can these two functions be merged into one to
// support both single and multiple UUIDs?

// Same here, can these two functions be merged into one?

func CassandraHotelToModel(cassandraHotel *CassandraHotel) (*model2.Hotel, error) {
	hotelUUID, err := uuid.FromBytes(cassandraHotel.HotelID.Bytes())
	if err != nil {
		return nil, err
	}

	pois, err := cassandraInfra.GocqlUUIDsToUUIDs(cassandraHotel.POIs)
	if err != nil {
		return nil, err
	}

	rooms, err := cassandraInfra.GocqlUUIDsToUUIDs(cassandraHotel.Rooms)
	if err != nil {
		return nil, err
	}

	return &model2.Hotel{
		HotelID: hotelUUID,
		Name:    cassandraHotel.Name,
		Address: model2.Address{
			Country:  cassandraHotel.Address.Country,
			City:     cassandraHotel.Address.City,
			Street:   cassandraHotel.Address.Street,
			Building: cassandraHotel.Address.Building,
			Entrance: cassandraHotel.Address.Entrance,
			ZipCode:  cassandraHotel.Address.ZipCode,
		},
		Description: cassandraHotel.Description,
		Phone:       cassandraHotel.Phone,
		POIs:        pois,
		Rooms:       rooms,
	}, nil
}

func ModelToCassandraHotel(hotel *model2.Hotel) (*CassandraHotel, error) {
	hotelID, err := gocql.ParseUUID(hotel.HotelID.String())
	if err != nil {
		return nil, err
	}

	pois, err := cassandraInfra.UUIDsToGocqlUUIDs(hotel.POIs)
	if err != nil {
		return nil, err
	}

	rooms, err := cassandraInfra.UUIDsToGocqlUUIDs(hotel.Rooms)
	if err != nil {
		return nil, err
	}

	return &CassandraHotel{
		HotelID: hotelID,
		Name:    hotel.Name,
		Address: CassandraAddress{
			Country:  hotel.Address.Country,
			City:     hotel.Address.City,
			Street:   hotel.Address.Street,
			Building: hotel.Address.Building,
			Entrance: hotel.Address.Entrance,
			ZipCode:  hotel.Address.ZipCode,
		},
		Description: hotel.Description,
		Phone:       hotel.Phone,
		POIs:        pois,
		Rooms:       rooms,
	}, nil
}
