package cassandra

import (
	"booking/internal/hotels/model"
	cassandraInfra "booking/internal/infra/cassandra"
	"log"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type HotelRepository struct {
	session *gocql.Session
}

func InitHotelRepository(session *gocql.Session) *HotelRepository {
	return &HotelRepository{session: session}
}

func (r *HotelRepository) GetByID(hotelUUID uuid.UUID) (*model.Hotel, error) {
	var hotel *model.Hotel

	var cassandraHotel CassandraHotel

	gocqlUUID, err := cassandraInfra.UUIDToGocqlUUID(hotelUUID)
	if err != nil {
		return nil, err
	}

	err = r.session.Query(
		`SELECT hotel_id,
				name,
				address,
				description,
				phone,
				pois,
				rooms
		FROM hotels WHERE hotel_id = ?`, gocqlUUID).Scan(
		&cassandraHotel.HotelID,
		&cassandraHotel.Name,
		&cassandraHotel.Address,
		&cassandraHotel.Description,
		&cassandraHotel.Phone,
		&cassandraHotel.POIs,
		&cassandraHotel.Rooms)
	if err != nil {
		log.Printf("Error fetching hotel by ID: %v", err)
		return nil, err
	}

	hotel, err = CassandraHotelToModel(&cassandraHotel)
	if err != nil {
		return nil, err
	}

	return hotel, nil
}

func (r *HotelRepository) Create(hotel *model.Hotel) error {
	cassandraHotel, err := ModelToCassandraHotel(hotel)
	if err != nil {
		return err
	}

	err = r.session.Query(`INSERT INTO hotels
	  (hotel_id, name, address, description, phone, pois, rooms)
	  VALUES (?, ?, ?, ?, ?, ?, ?)`,
		cassandraHotel.HotelID,
		cassandraHotel.Name,
		cassandraHotel.Address,
		cassandraHotel.Description,
		cassandraHotel.Phone,
		cassandraHotel.POIs,
		cassandraHotel.Rooms,
	).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *HotelRepository) Delete(hotelUUID uuid.UUID) error {
	gocqlUUID, err := cassandraInfra.UUIDToGocqlUUID(hotelUUID)
	if err != nil {
		return err
	}

	err = r.session.Query(
		`DELETE FROM hotels WHERE hotel_id = ?`,
		gocqlUUID).Exec()
	if err != nil {
		return err
	}
	return nil
}
