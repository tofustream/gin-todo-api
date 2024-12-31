package task

type TaskDTO struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	IsCompleted bool   `json:"is_completed"`
}

func NewTaskDTO(id string, description string, createdAt string, updatedAt string, isCompleted bool) TaskDTO {
	return TaskDTO{
		ID:          id,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		IsCompleted: isCompleted,
	}
}
