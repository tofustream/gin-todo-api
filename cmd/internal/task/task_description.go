package task

import "errors"

type TaskDescription struct {
	value string
}

func NewTaskDescription(value string) (TaskDescription, error) {
	if len(value) == 0 {
		return TaskDescription{}, errors.New("task description is empty")
	}
	if len(value) > 140 {
		return TaskDescription{}, errors.New("task description is too long")
	}

	return TaskDescription{value: value}, nil
}

func (t *TaskDescription) Value() string {
	return t.value
}
