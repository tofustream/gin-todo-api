package timestamp

import (
	"errors"
	"time"
)

type Timestamp struct {
	createdAt time.Time
	updatedAt time.Time
}

var (
	ErrEmptyCreatedAt   = errors.New("empty created at")
	ErrEmptyUpdatedAt   = errors.New("empty updated at")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
)

func NewTimestamp(createdAt time.Time, updatedAt time.Time) (Timestamp, error) {
	if createdAt.IsZero() {
		return Timestamp{}, ErrEmptyCreatedAt
	}
	if updatedAt.IsZero() {
		return Timestamp{}, ErrEmptyUpdatedAt
	}
	if createdAt.After(updatedAt) {
		return Timestamp{}, ErrInvalidTimestamp
	}

	return Timestamp{
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (t Timestamp) CreatedAt() time.Time {
	return t.createdAt
}

func (t Timestamp) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t Timestamp) Update() Timestamp {
	return Timestamp{
		createdAt: t.createdAt,
		updatedAt: time.Now(),
	}
}
