package task

import "errors"

type TaskID struct {
	value string
}

func NewTaskID(value string) (TaskID, error) {
	if value == "" {
		return TaskID{}, errors.New("task id is required")
	}
	return TaskID{value: value}, nil
}

func (t *TaskID) Value() string {
	return t.value
}
