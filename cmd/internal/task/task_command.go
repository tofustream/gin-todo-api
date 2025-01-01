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
	// TaskID をパースしてエラーをチェック
	parsedTaskID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, err
	}

	// TaskID を新規作成
	id, err := NewTaskID(parsedTaskID)
	if err != nil {
		return nil, err
	}

	// コマンドオブジェクトを返す
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

type MarkTaskAsIncompleteCommand struct {
	taskID TaskID
}

func NewMarkTaskAsIncompleteCommand(taskID string) (*MarkTaskAsIncompleteCommand, error) {
	// TaskID をパースしてエラーをチェック
	parsedTaskID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, err
	}

	// TaskID を新規作成
	id, err := NewTaskID(parsedTaskID)
	if err != nil {
		return nil, err
	}

	// コマンドオブジェクトを返す
	return &MarkTaskAsIncompleteCommand{taskID: id}, nil
}

func (c *MarkTaskAsIncompleteCommand) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindById(c.taskID)
	if err != nil {
		return TaskDTO{}, err
	}

	newTask := task.MarkAsIncomplete()
	return repository.Update(newTask)
}
