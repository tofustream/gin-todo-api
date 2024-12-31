package task

import (
	"github.com/google/uuid"
)

type ITaskApplicationService interface {
	FindAll() ([]TaskDTO, error)
	FindById(paramID string) (TaskDTO, error)
	Add(description string) (TaskDTO, error)
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

func (s *TaskApplicationService) FindById(paramID string) (TaskDTO, error) {
	parsedID, err := uuid.Parse(paramID)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return TaskDTO{}, err
	}
	id, err := NewTaskID(parsedID)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return TaskDTO{}, err
	}

	task, err := s.repository.FindById(id)
	if err != nil {
		return TaskDTO{}, err
	}

	return task, nil
}

func (s *TaskApplicationService) Add(description string) (TaskDTO, error) {
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

	task := NewTask(taskID, taskDescription)
	err = s.repository.Add(task)
	if err != nil {
		return TaskDTO{}, err
	}

	return createDTOFromTask(task), nil
}
