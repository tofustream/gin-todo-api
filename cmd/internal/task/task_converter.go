package task

func taskToDTO(task Task) TaskDTO {
	id := task.ID()
	description := task.Description()
	return NewTaskDTO(
		id.String(),
		description.Value(),
		task.CreatedAt().String(),
		task.UpdatedAt().String(),
		task.IsCompleted(),
		task.IsDeleted(),
	)
}
