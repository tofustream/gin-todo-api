package task

type ITaskApplicationService interface {
	FindAll() ([]TaskDTO, error)
}

type TaskApplicationService struct {
	repository ITaskRepository
}

func NewTaskApplicationService(repository ITaskRepository) *TaskApplicationService {
	return &TaskApplicationService{repository: repository}
}

func (s *TaskApplicationService) FindAll() ([]TaskDTO, error) {
	tasks, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
