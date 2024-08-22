package cassandra

import (
	"booking/internal/model"
	"log"

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

func UUIDToGocqlUUID(uuid uuid.UUID) (gocql.UUID, error) {
	gocqlUUID, err := gocql.ParseUUID(uuid.String())
	if err != nil {
		return gocql.UUID{}, err
	}
	return gocqlUUID, nil
}

// Can these two functions be merged into one to
// support both single and multiple UUIDs?

func UUIDsToGocqlUUIDs(uuids []uuid.UUID) ([]gocql.UUID, error) {
	result := make([]gocql.UUID, len(uuids))
	var err error
	for i, uuid := range uuids {
		result[i], err = UUIDToGocqlUUID(uuid)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func GocqlUUIDToUUID(gocqlUUID gocql.UUID) (uuid.UUID, error) {
	convertedUUID, err := uuid.Parse(gocqlUUID.String())
	if err != nil {
		log.Println("Error converting Cassandra UUID to UUID:", err)
		return uuid.Nil, err
	}
	return convertedUUID, nil
}

// Same here, can these two functions be merged into one?

func GocqlUUIDsToUUIDs(gocqlUUIDs []gocql.UUID) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, len(gocqlUUIDs))
	var err error
	for i, gocqlUUID := range gocqlUUIDs {
		result[i], err = GocqlUUIDToUUID(gocqlUUID)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func CassandraHotelToModel(cassandraHotel *CassandraHotel) (*model.Hotel, error) {
	hotelUUID, err := uuid.FromBytes(cassandraHotel.HotelID.Bytes())
	if err != nil {
		return nil, err
	}

	pois, err := GocqlUUIDsToUUIDs(cassandraHotel.POIs)
	if err != nil {
		return nil, err
	}

	rooms, err := GocqlUUIDsToUUIDs(cassandraHotel.Rooms)
	if err != nil {
		return nil, err
	}

	return &model.Hotel{
		HotelID: hotelUUID,
		Name:    cassandraHotel.Name,
		Address: model.Address{
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

func ModelToCassandraHotel(hotel *model.Hotel) (*CassandraHotel, error) {
	hotelID, err := gocql.ParseUUID(hotel.HotelID.String())
	if err != nil {
		return nil, err
	}

	pois, err := UUIDsToGocqlUUIDs(hotel.POIs)
	if err != nil {
		return nil, err
	}

	rooms, err := UUIDsToGocqlUUIDs(hotel.Rooms)
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
