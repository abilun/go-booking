package cassandra

import (
	"booking/internal/model"

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
	POIs        []CassandraPOI   `cql:"pois"`
	Rooms       []CassandraRoom  `cql:"rooms"`
}

type CassandraPOI struct {
	POIID       gocql.UUID `cql:"poi_id"`
	Name        string     `cql:"name"`
	Description string     `cql:"description"`
}

type CassandraRoom struct {
	RoomID gocql.UUID `cql:"room_id"`
	Number int        `cql:"number"`
}

func CassandraPOIToModel(cassandraPOI *CassandraPOI) model.POI {
	poiUUID, _ := uuid.FromBytes(cassandraPOI.POIID.Bytes())

	return model.POI{
		POIID:       poiUUID,
		Name:        cassandraPOI.Name,
		Description: cassandraPOI.Description,
	}
}

func CassandraRoomToModel(cassandraRoom *CassandraRoom) model.Room {
	roomUUID, _ := uuid.FromBytes(cassandraRoom.RoomID.Bytes())

	return model.Room{
		RoomID: roomUUID,
		Number: cassandraRoom.Number,
	}
}

func CassandraHotelToModel(cassandraHotel *CassandraHotel) (*model.Hotel, error) {
	hotelUUID, err := uuid.FromBytes(cassandraHotel.HotelID.Bytes())
	if err != nil {
		return nil, err
	}

	pois := make([]model.POI, len(cassandraHotel.POIs))
	for i, poi := range cassandraHotel.POIs {
		pois[i] = CassandraPOIToModel(&poi)
	}

	rooms := make([]model.Room, len(cassandraHotel.Rooms))
	for i, room := range cassandraHotel.Rooms {
		rooms[i] = CassandraRoomToModel(&room)
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

	pois := make([]CassandraPOI, len(hotel.POIs))
	for i, poi := range hotel.POIs {
		pois[i] = ModelToCassandraPOI(&poi)
	}

	rooms := make([]CassandraRoom, len(hotel.Rooms))
	for i, room := range hotel.Rooms {
		rooms[i] = ModelToCassandraRoom(&room)
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

func ModelToCassandraPOI(poi *model.POI) CassandraPOI {
	poiID, _ := gocql.ParseUUID(poi.POIID.String())

	return CassandraPOI{
		POIID:       poiID,
		Name:        poi.Name,
		Description: poi.Description,
	}
}

func ModelToCassandraRoom(room *model.Room) CassandraRoom {
	roomID, _ := gocql.ParseUUID(room.RoomID.String())

	return CassandraRoom{
		RoomID: roomID,
		Number: room.Number,
	}
}
