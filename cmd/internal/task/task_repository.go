package task

import "errors"

type ITaskRepository interface {
	FindAll() ([]TaskDTO, error)
	FindById(id TaskID) (TaskDTO, error)
}

type InMemoryTaskRepository struct {
	tasks map[TaskID]Task
}

func NewInMemoryTaskRepository(tasks map[TaskID]Task) ITaskRepository {
	return &InMemoryTaskRepository{
		tasks: tasks,
	}
}

func createDTOFromTask(task Task) TaskDTO {
	id := task.ID()
	description := task.Description()
	return NewTaskDTO(
		id.String(),
		description.Value(),
		task.CreatedAt().String(),
		task.UpdatedAt().String(),
		task.IsCompleted(),
	)
}

func (r *InMemoryTaskRepository) FindAll() ([]TaskDTO, error) {
	if len(r.tasks) == 0 {
		return nil, errors.New("no tasks found")
	}

	tasks := make([]TaskDTO, 0, len(r.tasks))
	for _, task := range r.tasks {
		dto := createDTOFromTask(task)
		tasks = append(tasks, dto)
	}
	return tasks, nil
}

func (r *InMemoryTaskRepository) FindById(id TaskID) (TaskDTO, error) {
	task, ok := r.tasks[id]
	if !ok {
		return TaskDTO{}, errors.New("task not found")
	}
	dto := createDTOFromTask(task)
	return dto, nil
}
