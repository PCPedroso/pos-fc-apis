package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(id string) (ID, error) {
	ret, err := uuid.Parse(id)
	return ID(ret), err
}
