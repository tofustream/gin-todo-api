package task

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tofustream/gin-todo-api/cmd/internal/account"
	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type ITaskRepository interface {
	FindAllByAccountID(accountID account.AccountID) ([]FindAllByAccountIDResponseDTO, error)
	FindTask(taskID TaskID, accountID account.AccountID) (Task, error)
	Add(task Task) error
	Update(task Task) (TaskDTO, error)
}

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) ITaskRepository {
	return &PostgresTaskRepository{
		db: db,
	}
}

func (r PostgresTaskRepository) FindAllByAccountID(
	accountID account.AccountID) ([]FindAllByAccountIDResponseDTO, error) {
	query := `
        SELECT id, description, created_at, updated_at, is_completed
        FROM tasks
        WHERE account_id = $1 AND is_deleted = false
    `
	rows, err := r.db.Query(query, accountID.Value())
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]FindAllByAccountIDResponseDTO, 0)
	for rows.Next() {
		var dto FindAllByAccountIDResponseDTO
		err = rows.Scan(
			&dto.ID,
			&dto.Description,
			&dto.CreatedAt,
			&dto.UpdatedAt,
			&dto.IsCompleted,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		tasks = append(tasks, dto)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return tasks, nil
}

func (r PostgresTaskRepository) FindTask(taskID TaskID, accountID account.AccountID) (Task, error) {
	query := "SELECT * FROM tasks WHERE id = $1 AND account_id = $2 AND is_deleted = false"
	rows := r.db.QueryRow(query, taskID.Value(), accountID.String())

	fetchedData, err := scanRowForFindTask(rows)
	if err != nil {
		return Task{}, fmt.Errorf("failed to scan row: %w", err)
	}

	task, err := createTaskFromFetchedData(fetchedData)
	if err != nil {
		return Task{}, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func scanRowForFindTask(rows *sql.Row) (map[string]interface{}, error) {
	var (
		fethcedTaskID          string
		fetchedDescription     string
		fetchedCreatedAt       string
		fetchedUpdatedAt       string
		fetchedCompletedStatus bool
		fetchedDeletedStatus   bool
		fetchedAccountID       string
	)

	err := rows.Scan(
		&fethcedTaskID,
		&fetchedDescription,
		&fetchedCreatedAt,
		&fetchedUpdatedAt,
		&fetchedCompletedStatus,
		&fetchedDeletedStatus,
		&fetchedAccountID,
	)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"taskID":          fethcedTaskID,
		"description":     fetchedDescription,
		"createdAt":       fetchedCreatedAt,
		"updatedAt":       fetchedUpdatedAt,
		"completedStatus": fetchedCompletedStatus,
		"deletedStatus":   fetchedDeletedStatus,
		"accountID":       fetchedAccountID,
	}, nil
}

func createTaskFromFetchedData(data map[string]interface{}) (Task, error) {
	taskIDValue, err := NewTaskIDFromString(data["taskID"].(string))
	if err != nil {
		return Task{}, fmt.Errorf("failed to create TaskID: %w", err)
	}
	descriptionValue, err := NewTaskDescription(data["description"].(string))
	if err != nil {
		return Task{}, fmt.Errorf("failed to create TaskDescription: %w", err)
	}
	createdAt, err := time.Parse(time.RFC3339, data["createdAt"].(string))
	if err != nil {
		return Task{}, fmt.Errorf("failed to parse createdAt: %w", err)
	}
	updatedAt, err := time.Parse(time.RFC3339, data["updatedAt"].(string))
	if err != nil {
		return Task{}, fmt.Errorf("failed to parse updatedAt: %w", err)
	}
	timeStamp, err := timestamp.NewTimestamp(createdAt, updatedAt)
	if err != nil {
		return Task{}, fmt.Errorf("failed to create Timestamp: %w", err)
	}
	accoutID, err := account.NewAccountIDFromString(data["accountID"].(string))
	if err != nil {
		return Task{}, fmt.Errorf("failed to create UserID: %w", err)
	}

	return NewTaskWithAllFields(
		taskIDValue,
		descriptionValue,
		timeStamp,
		data["completedStatus"].(bool),
		data["deletedStatus"].(bool),
		accoutID,
	), nil
}

func (r *PostgresTaskRepository) Add(task Task) error {
	taskIDValue := task.ID()
	descriptionValue := task.Description()
	query := `
        INSERT INTO tasks (id, description, created_at, updated_at, is_completed, is_deleted, account_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := r.db.Exec(
		query,
		taskIDValue.Value(),
		descriptionValue.Value(),
		task.CreatedAt().Format(time.RFC3339),
		task.UpdatedAt().Format(time.RFC3339),
		task.IsCompleted(),
		task.IsDeleted(),
		task.AccountID().Value(),
	)
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}
	return nil
}

func (r *PostgresTaskRepository) Update(task Task) (TaskDTO, error) {
	_, err := r.db.Exec(
		"UPDATE tasks SET description = $1, updated_at = $2, is_completed = $3, is_deleted = $4 WHERE id = $5",
		task.Description().Value(),
		task.UpdatedAt().Format(time.RFC3339),
		task.IsCompleted(),
		task.IsDeleted(),
		task.ID().Value())
	if err != nil {
		return TaskDTO{}, fmt.Errorf("failed to update task: %w", err)
	}
	return taskToDTO(task), nil
}
