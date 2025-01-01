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

	// TaskDescription を作成
	d, err := NewTaskDescription(description)
	if err != nil {
		return nil, err
	}

	// コマンドオブジェクトを返す
	return &UpdateTaskDescriptionCommand{
		taskID:      id,
		description: d,
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
