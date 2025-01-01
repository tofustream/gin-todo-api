package task

import (
	"github.com/google/uuid"
)

type ITaskCommand interface {
	Execute(repository ITaskRepository) (TaskDTO, error)
}

type UpdateTaskDescriptionCommand struct {
	taskID      TaskID
	description TaskDescription
}

func NewUpdateTaskDescriptionCommand(taskID string, description string) (*UpdateTaskDescriptionCommand, error) {
	parsedTaskID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, err
	}

	id, err := NewTaskID(parsedTaskID)
	if err != nil {
		return nil, err
	}

	desc, err := NewTaskDescription(description)
	if err != nil {
		return nil, err
	}

	return &UpdateTaskDescriptionCommand{
		taskID:      id,
		description: desc,
	}, nil
}

func (c *UpdateTaskDescriptionCommand) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindById(c.taskID)
	if err != nil {
		return TaskDTO{}, err
	}

	newTask := task.UpdateDescription(c.description)
	return repository.Update(newTask)
}

type MarkTaskAsCompleteCommand struct {
	taskID TaskID
}

func NewMarkTaskAsCompleteCommand(taskID string) (*MarkTaskAsCompleteCommand, error) {
	parsedTaskID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, err
	}

	id, err := NewTaskID(parsedTaskID)
	if err != nil {
		return nil, err
	}

	return &MarkTaskAsCompleteCommand{taskID: id}, nil
}

func (c *MarkTaskAsCompleteCommand) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindById(c.taskID)
	if err != nil {
		return TaskDTO{}, err
	}

	newTask := task.MarkAsComplete()
	return repository.Update(newTask)
}

type MarkTaskAsIncompleted struct {
	taskID TaskID
}

func NewMarkTaskAsIncompleteCommand(taskID string) (*MarkTaskAsIncompleted, error) {
	parsedTaskID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, err
	}

	// TaskID を新規作成
	id, err := NewTaskID(parsedTaskID)
	if err != nil {
		return nil, err
	}

	return &MarkTaskAsIncompleted{taskID: id}, nil
}

func (c *MarkTaskAsIncompleted) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindById(c.taskID)
	if err != nil {
		return TaskDTO{}, err
	}

	newTask := task.MarkAsIncomplete()
	return repository.Update(newTask)
}

type MarkAsDeletedCommand struct {
	taskID TaskID
}

func NewMarkAsDeletedCommand(taskID string) (*MarkAsDeletedCommand, error) {
	parsedTaskID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, err
	}

	id, err := NewTaskID(parsedTaskID)
	if err != nil {
		return nil, err
	}

	return &MarkAsDeletedCommand{taskID: id}, nil
}

func (c *MarkAsDeletedCommand) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindById(c.taskID)
	if err != nil {
		return TaskDTO{}, err
	}

	newTask := task.MarkAsDeleted()
	return repository.Update(newTask)
}
