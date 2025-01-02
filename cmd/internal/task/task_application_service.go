package task

import (
	"github.com/google/uuid"
	"github.com/tofustream/gin-todo-api/cmd/internal/account"
	"github.com/tofustream/gin-todo-api/cmd/internal/user"
)

type ITaskApplicationService interface {
	FindAll() ([]TaskDTO, error)
	FindAllByAccountID(accountID uuid.UUID) ([]TaskFindAllByAccountIDResponseDTO, error)
	FindById(paramID string) (TaskDTO, error)
	Register(description string, userID string) (TaskDTO, error)
	Update(command ITaskCommand) (TaskDTO, error)
}

type TaskApplicationService struct {
	repository ITaskRepository
}

func NewTaskApplicationService(repository ITaskRepository) ITaskApplicationService {
	return &TaskApplicationService{repository: repository}
}

func (s *TaskApplicationService) FindAll() ([]TaskDTO, error) {
	tasks, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
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
	parsedID, err := uuid.Parse(paramID)
	if err != nil {
		return TaskDTO{}, err
	}
	id, err := NewTaskID(parsedID)
	if err != nil {
		return TaskDTO{}, err
	}

	task, err := s.repository.FindById(id)
	if err != nil {
		return TaskDTO{}, err
	}

	return taskToDTO(task), nil
}

func (s *TaskApplicationService) Register(description string, userID string) (TaskDTO, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return TaskDTO{}, err
	}
	taskID, err := NewTaskID(newUUID)
	if err != nil {
		return TaskDTO{}, err
	}
	taskDescription, err := NewTaskDescription(description)
	if err != nil {
		return TaskDTO{}, err
	}

	userIDValue, err := user.NewUserIDFromString(userID)
	if err != nil {
		return TaskDTO{}, err
	}

	task := NewTask(taskID, taskDescription, userIDValue)
	err = s.repository.Add(task)
	if err != nil {
		return TaskDTO{}, err
	}

	return taskToDTO(task), nil
}

func (s *TaskApplicationService) Update(command ITaskCommand) (TaskDTO, error) {
	return command.Execute(s.repository)
}
