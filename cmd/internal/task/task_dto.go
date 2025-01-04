package task

type TaskDTO struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	IsCompleted bool   `json:"is_completed"`
	IsDeleted   bool   `json:"is_deleted"`
	AccountID   string `json:"account_id"`
}
