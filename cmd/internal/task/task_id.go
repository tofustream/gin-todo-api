package task

import "errors"

type TaskID struct {
	value uint
}

func NewTaskID(value uint) (TaskID, error) {
	if value == 0 {
		return TaskID{}, errors.New("task ID is zero")
	}
	return TaskID{value: value}, nil
}

func (t *TaskID) Value() uint {
	return t.value
}
