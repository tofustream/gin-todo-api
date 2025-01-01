package task

import (
	"github.com/google/uuid"
)

type ITaskApplicationService interface {
	FindAll() ([]TaskDTO, error)
	FindById(paramID string) (TaskDTO, error)
	Register(description string) (TaskDTO, error)
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

func (s *TaskApplicationService) Register(description string) (TaskDTO, error) {
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

	return taskToDTO(task), nil
}

func (s *TaskApplicationService) Update(command ITaskCommand) (TaskDTO, error) {
	return command.Execute(s.repository)
}
