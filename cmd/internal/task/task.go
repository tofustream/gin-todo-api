package task

import "time"

type Task struct {
	id          TaskID
	description TaskDescription
	createdAt   time.Time
	updatedAt   time.Time
	isCompleted bool
}

func NewTask(id TaskID, description TaskDescription) Task {
	now := time.Now()
	return Task{id: id, description: description, createdAt: now, updatedAt: now, isCompleted: false}
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

func (t *Task) MarkAsComplete() Task {
	return Task{
		id:          t.id,
		description: t.description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isCompleted: true,
	}
}

func (t *Task) MarkAsIncomplete() Task {
	return Task{
		id:          t.id,
		description: t.description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isCompleted: false,
	}
}

func (t *Task) UpdateDescription(description TaskDescription) Task {
	return Task{
		id:          t.id,
		description: description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isCompleted: t.isCompleted,
	}
}
