package task

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tofustream/gin-todo-api/cmd/internal/account"
	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type ITaskRepository interface {
	FindAllTasksByAccountID(accountID account.AccountID) ([]Task, error)
	FindTask(taskID TaskID, accountID account.AccountID) (*Task, error)
	AddTask(task Task) error
	UpdateTask(task Task) error
}

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) ITaskRepository {
	return &PostgresTaskRepository{
		db: db,
	}
}

func (r PostgresTaskRepository) FindAllTasksByAccountID(
	accountID account.AccountID,
) ([]Task, error) {
	query := `
		SELECT * FROM tasks
		WHERE account_id = $1 AND is_deleted = false
	`
	rows, err := r.db.Query(query, accountID.Value())
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var (
			fethcedTaskIDValue      string
			fetchedDescriptionValue string
			fetchedCreatedAtValue   time.Time
			fetchedUpdatedAtValue   time.Time
			isCompleted             bool
			isDeleted               bool
			fetchedAccountIDValue   string
		)
		err = rows.Scan(
			&fethcedTaskIDValue,
			&fetchedDescriptionValue,
			&fetchedCreatedAtValue,
			&fetchedUpdatedAtValue,
			&isCompleted,
			&isDeleted,
			&fetchedAccountIDValue,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		fetchedTaskID, err := NewTaskIDFromString(fethcedTaskIDValue)
		if err != nil {
			return nil, fmt.Errorf("failed to create TaskID: %w", err)
		}
		fetchedDescription, err := NewTaskDescription(fetchedDescriptionValue)
		if err != nil {
			return nil, fmt.Errorf("failed to create TaskDescription: %w", err)
		}
		fetchedTimestamp, err := timestamp.NewTimestamp(fetchedCreatedAtValue, fetchedUpdatedAtValue)
		if err != nil {
			return nil, fmt.Errorf("failed to create Timestamp: %w", err)
		}
		fetchedAccountID, err := account.NewAccountIDFromString(fetchedAccountIDValue)
		if err != nil {
			return nil, fmt.Errorf("failed to create AccountID: %w", err)
		}
		task := NewTaskWithAllFields(
			fetchedTaskID,
			fetchedDescription,
			fetchedTimestamp,
			isCompleted,
			isDeleted,
			fetchedAccountID,
		)

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return tasks, nil
}

func (r PostgresTaskRepository) FindTask(taskID TaskID, accountID account.AccountID) (*Task, error) {
	query := "SELECT * FROM tasks WHERE id = $1 AND account_id = $2 AND is_deleted = false"
	rows := r.db.QueryRow(query, taskID.Value(), accountID.String())

	var (
		fethcedTaskIDValue      string
		fetchedDescriptionValue string
		fetchedCreatedAtValue   time.Time
		fetchedUpdatedAtValue   time.Time
		isCompleted             bool
		isDeleted               bool
		fetchedAccountIDValue   string
	)
	err := rows.Scan(
		&fethcedTaskIDValue,
		&fetchedDescriptionValue,
		&fetchedCreatedAtValue,
		&fetchedUpdatedAtValue,
		&isCompleted,
		&isDeleted,
		&fetchedAccountIDValue,
	)
	if err != nil {
		return nil, err
	}

	fetchedTaskID, err := NewTaskIDFromString(fethcedTaskIDValue)
	if err != nil {
		return nil, fmt.Errorf("failed to create TaskID: %w", err)
	}
	fetchedDescription, err := NewTaskDescription(fetchedDescriptionValue)
	if err != nil {
		return nil, fmt.Errorf("failed to create TaskDescription: %w", err)
	}
	fetchedTimestamp, err := timestamp.NewTimestamp(fetchedCreatedAtValue, fetchedUpdatedAtValue)
	if err != nil {
		return nil, fmt.Errorf("failed to create Timestamp: %w", err)
	}
	fetchedAccountID, err := account.NewAccountIDFromString(fetchedAccountIDValue)
	if err != nil {
		return nil, fmt.Errorf("failed to create AccountID: %w", err)
	}
	task := NewTaskWithAllFields(
		fetchedTaskID,
		fetchedDescription,
		fetchedTimestamp,
		isCompleted,
		isDeleted,
		fetchedAccountID,
	)

	return &task, nil
}

func (r PostgresTaskRepository) AddTask(task Task) error {
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

func (r PostgresTaskRepository) UpdateTask(task Task) error {
	_, err := r.db.Exec(
		"UPDATE tasks SET description = $1, updated_at = $2, is_completed = $3, is_deleted = $4 WHERE id = $5",
		task.Description().Value(),
		task.UpdatedAt().Format(time.RFC3339),
		task.IsCompleted(),
		task.IsDeleted(),
		task.ID().Value(),
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}
