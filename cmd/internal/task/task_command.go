package task

import "github.com/tofustream/gin-todo-api/cmd/internal/account"

type ITaskCommand interface {
	Execute(repository ITaskRepository) (*TaskDTO, error)
}

type UpdateTaskDescriptionCommand struct {
	taskID      TaskID
	description TaskDescription
	accountID   account.AccountID
}

func NewUpdateTaskDescriptionCommand(taskID string, description string, accountID string) (*UpdateTaskDescriptionCommand, error) {
	taskIDValue, err := NewTaskIDFromString(taskID)
	if err != nil {
		return nil, err
	}

	descriptionValue, err := NewTaskDescription(description)
	if err != nil {
		return nil, err
	}

	accountIDValue, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	return &UpdateTaskDescriptionCommand{
		taskID:      taskIDValue,
		description: descriptionValue,
		accountID:   accountIDValue,
	}, nil
}

func (c UpdateTaskDescriptionCommand) Execute(repository ITaskRepository) (*TaskDTO, error) {
	task, err := repository.FindTask(c.taskID, c.accountID)
	if err != nil {
		return nil, err
	}
	newTask := task.UpdateDescription(c.description)
	return repository.UpdateTask(newTask)
}

type MarkAsDeletedCommand struct {
	taskID    TaskID
	accountID account.AccountID
}

func NewMarkAsDeletedCommand(taskID string, accountID string) (*MarkAsDeletedCommand, error) {
	taskIDValue, err := NewTaskIDFromString(taskID)
	if err != nil {
		return nil, err
	}
	accountIDValue, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	return &MarkAsDeletedCommand{
		taskID:    taskIDValue,
		accountID: accountIDValue,
	}, nil
}

func (c *MarkAsDeletedCommand) Execute(repository ITaskRepository) (*TaskDTO, error) {
	task, err := repository.FindTask(c.taskID, c.accountID)
	if err != nil {
		return nil, err
	}

	newTask := task.MarkAsDeleted()
	return repository.UpdateTask(newTask)
}

type UpdateTaskStatusCommand struct {
	taskID      TaskID
	isCompleted bool
	accountID   account.AccountID
}

func NewUpdateTaskStatusCommand(taskID string, isCompleted bool, accountID string) (ITaskCommand, error) {
	taskIDInstance, err := NewTaskIDFromString(taskID)
	if err != nil {
		return nil, err
	}

	accountIDInstance, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	return &UpdateTaskStatusCommand{
		taskID:      taskIDInstance,
		isCompleted: isCompleted,
		accountID:   accountIDInstance,
	}, nil
}

func (c UpdateTaskStatusCommand) Execute(repository ITaskRepository) (*TaskDTO, error) {
	task, err := repository.FindTask(c.taskID, c.accountID)
	if err != nil {
		return nil, err
	}

	if c.isCompleted {
		newTask := task.MarkAsComplete()
		return repository.UpdateTask(newTask)
	}

	newTask := task.MarkAsIncomplete()
	return repository.UpdateTask(newTask)
}
