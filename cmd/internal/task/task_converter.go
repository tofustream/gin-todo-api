package task

// is_deleteは自明なので省略
func taskToDTO(task Task) *TaskDTO {
	dto := TaskDTO{
		ID:          task.ID().String(),
		Description: task.Description().Value(),
		CreatedAt:   task.CreatedAt().String(),
		UpdatedAt:   task.UpdatedAt().String(),
		IsCompleted: task.IsCompleted(),
		AccountID:   task.AccountID().String(),
	}
	return &dto
}
