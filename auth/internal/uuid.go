package auth

import (
	"fmt"

	"github.com/google/uuid"
)

type UUID struct {
	uuid uuid.UUID
}

func NewUUID(uuid uuid.UUID) UUID {
	return UUID{uuid: uuid}
}

func NewUUIDByByte(b []byte) (UUID, error) {
	uid, err := uuid.FromBytes(b)
	if err != nil {
		return UUID{}, err
	}
	return UUID{uuid: uid}, nil
}

func NewUUIDByString(id string) (UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return UUID{}, fmt.Errorf("Check the uuid format")
	}
	return UUID{uuid: uid}, nil
}

func (u UUID) String() string {
	return u.uuid.String()
}
