package task

import (
	"time"

	"github.com/tofustream/gin-todo-api/cmd/internal/user"
)

type Task struct {
	id          TaskID
	description TaskDescription
	createdAt   time.Time
	updatedAt   time.Time
	isCompleted bool
	isDeleted   bool
	userID      user.UserID
}

func NewTask(id TaskID, description TaskDescription, userID user.UserID) Task {
	now := time.Now()
	return Task{
		id:          id,
		description: description,
		createdAt:   now,
		updatedAt:   now,
		isCompleted: false,
		isDeleted:   false,
		userID:      userID,
	}
}

func NewTaskWithAllFields(
	id TaskID,
	description TaskDescription,
	createdAt, updatedAt time.Time,
	isCompleted, isDeleted bool,
	userID user.UserID) Task {
	return Task{
		id:          id,
		description: description,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		isCompleted: isCompleted,
		isDeleted:   isDeleted,
		userID:      userID,
	}
}

func (t *Task) ID() TaskID {
	return t.id
}

func (t *Task) Description() TaskDescription {
	return t.description
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) IsCompleted() bool {
	return t.isCompleted
}

func (t *Task) IsDeleted() bool {
	return t.isDeleted
}

func (t *Task) UserID() user.UserID {
	return t.userID
}

func (t *Task) MarkAsComplete() Task {
	return Task{
		id:          t.id,
		description: t.description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isCompleted: true,
		isDeleted:   t.isDeleted,
		userID:      t.userID,
	}
}

func (t *Task) MarkAsIncomplete() Task {
	return Task{
		id:          t.id,
		description: t.description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isCompleted: false,
		isDeleted:   t.isDeleted,
		userID:      t.userID,
	}
}

func (t *Task) UpdateDescription(description TaskDescription) Task {
	return Task{
		id:          t.id,
		description: description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isCompleted: t.isCompleted,
		isDeleted:   t.isDeleted,
		userID:      t.userID,
	}
}

func (t *Task) MarkAsDeleted() Task {
	return Task{
		id:          t.id,
		description: t.description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isCompleted: t.isCompleted,
		isDeleted:   true,
		userID:      t.userID,
	}
}
