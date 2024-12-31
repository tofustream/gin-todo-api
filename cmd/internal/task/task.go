package task

import "time"

type Task struct {
	id          TaskID
	description TaskDescription
	createdAt   time.Time
	updatedAt   time.Time
	isDeleted   bool
}

func NewTask(id TaskID, description TaskDescription) Task {
	now := time.Now()
	return Task{id: id, description: description, createdAt: now, updatedAt: now, isDeleted: false}
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

func (t *Task) IsDeleted() bool {
	return t.isDeleted
}

func (t *Task) Delete() Task {
	return Task{
		id:          t.id,
		description: t.description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isDeleted:   true,
	}
}

func (t *Task) Restore() Task {
	return Task{
		id:          t.id,
		description: t.description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isDeleted:   false,
	}
}

func (t *Task) ChangeDescription(description TaskDescription) Task {
	return Task{
		id:          t.id,
		description: description,
		createdAt:   t.createdAt,
		updatedAt:   time.Now(),
		isDeleted:   t.isDeleted,
	}
}
