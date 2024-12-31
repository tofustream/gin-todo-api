package task

type ITaskRepository interface {
	FindAll() ([]Task, error)
}

type InMemoryTaskRepository struct {
	tasks map[TaskID]Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[TaskID]Task),
	}
}

func (r *InMemoryTaskRepository) FindAll() ([]Task, error) {
	tasks := make([]Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
