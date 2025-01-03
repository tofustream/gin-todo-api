package task

import "github.com/tofustream/gin-todo-api/cmd/internal/account"

type ITaskCommand interface {
	Execute(repository ITaskRepository) (TaskDTO, error)
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

func (c UpdateTaskDescriptionCommand) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindTask(c.taskID, c.accountID)
	if err != nil {
		return TaskDTO{}, err
	}
	newTask := task.UpdateDescription(c.description)
	return repository.Update(newTask)
}

type MarkTaskAsCompleteCommand struct {
	taskID    TaskID
	accountID account.AccountID
}

func NewMarkTaskAsCompleteCommand(taskID string, accountID string) (*MarkTaskAsCompleteCommand, error) {
	taskIDValue, err := NewTaskIDFromString(taskID)
	if err != nil {
		return nil, err
	}

	accountIDValue, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	return &MarkTaskAsCompleteCommand{
		taskID:    taskIDValue,
		accountID: accountIDValue,
	}, nil
}

func (c *MarkTaskAsCompleteCommand) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindTask(c.taskID, c.accountID)
	if err != nil {
		return TaskDTO{}, err
	}

	newTask := task.MarkAsComplete()
	return repository.Update(newTask)
}

type MarkTaskAsIncompleted struct {
	taskID    TaskID
	accountID account.AccountID
}

func NewMarkTaskAsIncompleteCommand(taskID string, accountID string) (*MarkTaskAsIncompleted, error) {
	taskIDValue, err := NewTaskIDFromString(taskID)
	if err != nil {
		return nil, err
	}
	accountIDValue, err := account.NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	return &MarkTaskAsIncompleted{
		taskID:    taskIDValue,
		accountID: accountIDValue,
	}, nil
}

func (c *MarkTaskAsIncompleted) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindTask(c.taskID, c.accountID)
	if err != nil {
		return TaskDTO{}, err
	}

	newTask := task.MarkAsIncomplete()
	return repository.Update(newTask)
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

func (c *MarkAsDeletedCommand) Execute(repository ITaskRepository) (TaskDTO, error) {
	task, err := repository.FindTask(c.taskID, c.accountID)
	if err != nil {
		return TaskDTO{}, err
	}

	newTask := task.MarkAsDeleted()
	return repository.Update(newTask)
}
