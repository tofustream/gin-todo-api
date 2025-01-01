package task

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

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

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) ITaskRepository {
	return &PostgresTaskRepository{
		db: db,
	}
}

func (r *PostgresTaskRepository) FindAll() ([]TaskDTO, error) {
	rows, err := r.db.Query("SELECT * FROM tasks WHERE is_deleted = false")
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]TaskDTO, 0)
	for rows.Next() {
		var id string
		var description string
		var createdAt string
		var updatedAt string
		var isCompleted bool
		var isDeleted bool
		err = rows.Scan(&id, &description, &createdAt, &updatedAt, &isCompleted, &isDeleted)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		tasks = append(tasks, TaskDTO{
			ID:          id,
			Description: description,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			IsCompleted: isCompleted,
			IsDeleted:   isDeleted,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return tasks, nil
}

func (r *PostgresTaskRepository) FindById(id TaskID) (Task, error) {
	var task Task
	var fetchedID string
	var fetchedDescription string
	var fetchedCreatedAt string
	var fetchedUpdatedAt string
	var fetchedCompletedStatus bool
	var fetchedDeletedStatus bool

	row := r.db.QueryRow("SELECT * FROM tasks WHERE id = $1 AND is_deleted = false", id.Value())
	err := row.Scan(
		&fetchedID,
		&fetchedDescription,
		&fetchedCreatedAt,
		&fetchedUpdatedAt,
		&fetchedCompletedStatus,
		&fetchedDeletedStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, errors.New("task not found")
		}
		return Task{}, fmt.Errorf("failed to scan row: %w", err)
	}

	parsedID, err := uuid.Parse(fetchedID)
	if err != nil {
		return Task{}, fmt.Errorf("failed to parse fetchedID: %w", err)
	}
	taskID, err := NewTaskID(parsedID)
	if err != nil {
		return Task{}, fmt.Errorf("failed to create TaskID: %w", err)
	}
	desc, err := NewTaskDescription(fetchedDescription)
	if err != nil {
		return Task{}, fmt.Errorf("failed to create TaskDescription: %w", err)
	}
	createdAt, err := time.Parse(time.RFC3339, fetchedCreatedAt)
	if err != nil {
		return Task{}, fmt.Errorf("failed to parse createdAt: %w", err)
	}
	updatedAt, err := time.Parse(time.RFC3339, fetchedUpdatedAt)
	if err != nil {
		return Task{}, fmt.Errorf("failed to parse updatedAt: %w", err)
	}

	task = NewTaskWithAllFields(
		taskID,
		desc,
		createdAt,
		updatedAt,
		fetchedCompletedStatus,
		fetchedDeletedStatus)
	return task, nil
}

func (r *PostgresTaskRepository) Add(task Task) error {
	id := task.ID()
	desc := task.Description()
	_, err := r.db.Exec("INSERT INTO tasks (id, description, created_at, updated_at, is_completed, is_deleted) VALUES ($1, $2, $3, $4, $5, $6)",
		id.Value(),
		desc.Value(),
		task.CreatedAt().Format(time.RFC3339),
		task.UpdatedAt().Format(time.RFC3339),
		task.IsCompleted(),
		task.IsDeleted())
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}
	return nil
}

func (r *PostgresTaskRepository) Update(task Task) (TaskDTO, error) {
	id := task.ID()
	desc := task.Description()
	_, err := r.db.Exec("UPDATE tasks SET description = $1, updated_at = $2, is_completed = $3, is_deleted = $4 WHERE id = $5",
		desc.Value(),
		task.UpdatedAt().Format(time.RFC3339),
		task.IsCompleted(),
		task.IsDeleted(),
		id.Value())
	if err != nil {
		return TaskDTO{}, fmt.Errorf("failed to update task: %w", err)
	}
	return taskToDTO(task), nil
}
