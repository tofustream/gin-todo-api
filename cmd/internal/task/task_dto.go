package task

import "time"

type TaskDTO struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	IsCompleted bool   `json:"is_completed"`
	IsDeleted   bool   `json:"is_deleted"`
}

func NewTaskDTO(
	id string,
	description string,
	createdAt string,
	updatedAt string,
	isCompleted bool,
	isDeleted bool) TaskDTO {
	return TaskDTO{
		ID:          id,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		IsCompleted: isCompleted,
		IsDeleted:   isDeleted,
	}
}

type TaskFindAllByAccountIDResponseDTO struct {
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsCompleted bool      `json:"is_completed"`
}
