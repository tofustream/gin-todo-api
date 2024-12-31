package task

type ITaskApplicationService interface {
	FindAll() ([]TaskDTO, error)
	FindById(id TaskID) (TaskDTO, error)
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

func (s *TaskApplicationService) FindById(id TaskID) (TaskDTO, error) {
	task, err := s.repository.FindById(id)
	if err != nil {
		return TaskDTO{}, err
	}

	return task, nil
}
