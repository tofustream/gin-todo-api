package task

import "errors"

type ITaskRepository interface {
	FindAll() ([]TaskDTO, error)
	FindById(id TaskID) (Task, error)
	Add(task Task) error
	Update(task Task) (TaskDTO, error)
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
		dto := taskToDTO(task)
		tasks = append(tasks, dto)
	}
	return tasks, nil
}

func (r *InMemoryTaskRepository) FindById(id TaskID) (Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryTaskRepository) Add(task Task) error {
	if _, exists := r.tasks[task.ID()]; exists {
		return errors.New("task already exists")
	}
	r.tasks[task.ID()] = task
	return nil
}

func (r *InMemoryTaskRepository) Update(task Task) (TaskDTO, error) {
	if _, exists := r.tasks[task.ID()]; !exists {
		return TaskDTO{}, errors.New("task not found")
	}
	r.tasks[task.ID()] = task
	return taskToDTO(r.tasks[task.ID()]), nil
}
