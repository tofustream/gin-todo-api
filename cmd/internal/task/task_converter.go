package task

func taskToDTO(task Task) *TaskDTO {
	dto := TaskDTO{
		ID:          task.ID().String(),
		Description: task.Description().Value(),
		CreatedAt:   task.CreatedAt().String(),
		UpdatedAt:   task.UpdatedAt().String(),
		IsCompleted: task.IsCompleted(),
		IsDeleted:   task.IsDeleted(),
		AccountID:   task.AccountID().String(),
	}
	return &dto
}
