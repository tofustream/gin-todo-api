package account

import (
	"errors"

	"github.com/google/uuid"
)

type AccountID struct {
	value uuid.UUID
}

func NewAccountIDFromUUID(value uuid.UUID) (AccountID, error) {
	if value == uuid.Nil {
		return AccountID{}, errors.New("invalid account ID")
	}
	return AccountID{value: value}, nil
}

func NewAccountIDFromString(value string) (AccountID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return AccountID{}, err
	}
	return NewAccountIDFromUUID(id)
}

func (u AccountID) Value() uuid.UUID {
	return u.value
}

func (u AccountID) Equals(other AccountID) bool {
	return u.value == other.value
}

func (u AccountID) String() string {
	return u.value.String()
}
