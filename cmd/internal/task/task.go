package task

import (
	"time"

	"github.com/tofustream/gin-todo-api/cmd/internal/account"
	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type Task struct {
	id          TaskID
	description TaskDescription
	// createdAt   time.Time
	// updatedAt   time.Time
	timeStamp   timestamp.Timestamp
	isCompleted bool
	isDeleted   bool
	accountID   account.AccountID
}

func NewTask(id TaskID, description TaskDescription, accountID account.AccountID) Task {
	now := time.Now()
	timeStamp, _ := timestamp.NewTimestamp(now, now)
	return Task{
		id:          id,
		description: description,
		timeStamp:   timeStamp,
		isCompleted: false,
		isDeleted:   false,
		accountID:   accountID,
	}
}

func NewTaskWithAllFields(
	id TaskID,
	description TaskDescription,
	timeStamp timestamp.Timestamp,
	isCompleted, isDeleted bool,
	accountID account.AccountID,
) Task {
	return Task{
		id:          id,
		description: description,
		timeStamp:   timeStamp,
		isCompleted: isCompleted,
		isDeleted:   isDeleted,
		accountID:   accountID,
	}
}

func (t Task) ID() TaskID {
	return t.id
}

func (t Task) Description() TaskDescription {
	return t.description
}

func (t Task) CreatedAt() time.Time {
	return t.timeStamp.CreatedAt()
}

func (t Task) UpdatedAt() time.Time {
	return t.timeStamp.UpdatedAt()
}

func (t Task) IsCompleted() bool {
	return t.isCompleted
}

func (t Task) IsDeleted() bool {
	return t.isDeleted
}

func (t Task) AccountID() account.AccountID {
	return t.accountID
}

func (t Task) MarkAsComplete() Task {
	newTimestamp := t.timeStamp.Update()
	return Task{
		id:          t.id,
		description: t.description,
		timeStamp:   newTimestamp,
		isCompleted: true,
		isDeleted:   t.isDeleted,
		accountID:   t.accountID,
	}
}

func (t Task) MarkAsIncomplete() Task {
	newTimestamp := t.timeStamp.Update()
	return Task{
		id:          t.id,
		description: t.description,
		timeStamp:   newTimestamp,
		isCompleted: false,
		isDeleted:   t.isDeleted,
		accountID:   t.accountID,
	}
}

func (t Task) UpdateDescription(description TaskDescription) Task {
	newTimestamp := t.timeStamp.Update()
	return Task{
		id:          t.id,
		description: description,
		timeStamp:   newTimestamp,
		isCompleted: t.isCompleted,
		isDeleted:   t.isDeleted,
		accountID:   t.accountID,
	}
}

func (t Task) MarkAsDeleted() Task {
	newTimestamp := t.timeStamp.Update()
	return Task{
		id:          t.id,
		description: t.description,
		timeStamp:   newTimestamp,
		isCompleted: t.isCompleted,
		isDeleted:   true,
		accountID:   t.accountID,
	}
}
