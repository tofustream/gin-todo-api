package user

import (
	"errors"

	"github.com/google/uuid"
)

type UserID struct {
	value uuid.UUID
}

func NewUserID(value uuid.UUID) (UserID, error) {
	if value == uuid.Nil {
		return UserID{}, errors.New("invalid user ID")
	}
	return UserID{value: value}, nil
}

func (u UserID) Value() uuid.UUID {
	return u.value
}

func (u UserID) Equals(other UserID) bool {
	return u.value == other.value
}