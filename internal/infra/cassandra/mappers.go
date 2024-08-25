package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"log"
)

func UUIDToGocqlUUID(uuid uuid.UUID) (gocql.UUID, error) {
	gocqlUUID, err := gocql.ParseUUID(uuid.String())
	if err != nil {
		return gocql.UUID{}, err
	}
	return gocqlUUID, nil
}

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
