package task

import (
	"errors"

	"github.com/google/uuid"
)

type TaskID struct {
	value uuid.UUID
}

func NewTaskIDFromUUID(value uuid.UUID) (TaskID, error) {
	if value == uuid.Nil {
		return TaskID{}, errors.New("task id cannot be nil")
	}
	return TaskID{value: value}, nil
}

func NewTaskIDFromString(value string) (TaskID, error) {
	parsed, err := uuid.Parse(value)
	if err != nil {
		return TaskID{}, err
	}
	return NewTaskIDFromUUID(parsed)
}

func (t TaskID) Value() uuid.UUID {
	return t.value
}

func (t TaskID) Equals(other TaskID) bool {
	return t.value == other.value
}

func (t TaskID) String() string {
	return t.value.String()
}
