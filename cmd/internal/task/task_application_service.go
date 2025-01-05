package task

import (
	"github.com/google/uuid"
	"github.com/tofustream/gin-todo-api/cmd/internal/account"
)

type ITaskApplicationService interface {
	FindAllTasksByAccountID(accountID string) ([]TaskDTO, error)
	FindTask(taskID string, accountID string) (*TaskDTO, error)
	CreateTask(description string, accountID string) error
	UpdateTask(command ITaskCommand) error
}

type TaskApplicationService struct {
	repository ITaskRepository
}

func NewTaskApplicationService(repository ITaskRepository) ITaskApplicationService {
	return &TaskApplicationService{repository: repository}
}

func (s TaskApplicationService) FindAllTasksByAccountID(accountID string) ([]TaskDTO, error) {
	accountIDInstance, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}
	dtos, err := s.repository.FindAllTasksByAccountID(accountIDInstance)
	if err != nil {
		return nil, err
	}
	return dtos, nil
}

func (s TaskApplicationService) FindTask(taskID string, accountID string) (*TaskDTO, error) {
	taskIDInstance, err := NewTaskIDFromString(taskID)
	if err != nil {
		return nil, err
	}
	accountIDInstance, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	task, err := s.repository.FindTask(taskIDInstance, accountIDInstance)
	if err != nil {
		return nil, err
	}

	return taskToDTO(task), nil
}

func (s TaskApplicationService) CreateTask(description string, accountID string) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	taskID, err := NewTaskIDFromUUID(newUUID)
	if err != nil {
		return err
	}
	taskDescription, err := NewTaskDescription(description)
	if err != nil {
		return err
	}

	accountIDInstance, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return err
	}

	task := NewTask(taskID, taskDescription, accountIDInstance)
	return s.repository.AddTask(task)
}

func (s TaskApplicationService) UpdateTask(command ITaskCommand) error {
	return command.Execute(s.repository)
}
