package task

import "errors"

type ITaskRepository interface {
	FindAll() ([]TaskDTO, error)
}

type InMemoryTaskRepository struct {
	tasks map[TaskID]Task
}

func NewInMemoryTaskRepository(tasks map[TaskID]Task) ITaskRepository {
	return &InMemoryTaskRepository{
		tasks: tasks,
	}
}

func (r *InMemoryTaskRepository) FindAll() ([]TaskDTO, error) {
	if len(r.tasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	tasks := make([]TaskDTO, 0, len(r.tasks))
	for _, task := range r.tasks {
		id := task.ID()
		description := task.Description()
		tasks = append(
			tasks,
			NewTaskDTO(
				id.String(),
				description.Value(),
				task.CreatedAt().String(),
				task.UpdatedAt().String(),
				task.IsCompleted(),
			),
		)
	}
	return tasks, nil
}
