package task

import (
	"log"

	"github.com/google/uuid"
	"github.com/tofustream/gin-todo-api/cmd/internal/account"
)

type ITaskApplicationService interface {
	FindAllByAccountID(accountID uuid.UUID) ([]TaskFindAllByAccountIDResponseDTO, error)
	FindById(paramID string) (TaskDTO, error)
	CreateTask(description string, accountID string) error
	Update(command ITaskCommand) (TaskDTO, error)
}

type TaskApplicationService struct {
	repository ITaskRepository
}

func NewTaskApplicationService(repository ITaskRepository) ITaskApplicationService {
	return &TaskApplicationService{repository: repository}
}

func (s *TaskApplicationService) FindAllByAccountID(accountID uuid.UUID) ([]TaskFindAllByAccountIDResponseDTO, error) {
	accountIDValue, err := account.NewAccountIDFromUUID(accountID)
	if err != nil {
		return nil, err
	}
	dtos, err := s.repository.FindAllByAccountID(accountIDValue)
	if err != nil {
		return nil, err
	}
	return dtos, nil
}

func (s *TaskApplicationService) FindById(paramID string) (TaskDTO, error) {
	id, err := NewTaskIDFromString(paramID)
	if err != nil {
		return TaskDTO{}, err
	}

	task, err := s.repository.FindById(id)
	if err != nil {
		return TaskDTO{}, err
	}

	return taskToDTO(task), nil
}

func (s *TaskApplicationService) CreateTask(description string, accountID string) error {
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

	accountIDValue, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return err
	}

	task := NewTask(taskID, taskDescription, accountIDValue)
	log.Printf("task: %v", task)
	return s.repository.Add(task)
}

func (s *TaskApplicationService) Update(command ITaskCommand) (TaskDTO, error) {
	return command.Execute(s.repository)
}
