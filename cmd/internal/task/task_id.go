package task

import (
	"errors"

	"github.com/google/uuid"
)

type TaskID struct {
	value uuid.UUID
}

func NewTaskID(value uuid.UUID) (TaskID, error) {
	if value == uuid.Nil {
		return TaskID{}, errors.New("task id cannot be nil")
	}
	return TaskID{value: value}, nil
}

func (t *TaskID) Value() uuid.UUID {
	return t.value
}

func (t *TaskID) Equals(other TaskID) bool {
	return t.value == other.value
}

func (t *TaskID) String() string {
	return t.value.String()
}
