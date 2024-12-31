package task_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/tofustream/gin-todo-api/cmd/internal/task"
)

func TestNewTask(t *testing.T) {
	validUUID, _ := uuid.NewRandom()
	id, _ := task.NewTaskID(validUUID)
	description, _ := task.NewTaskDescription("Test task")
	task := task.NewTask(id, description)

	if task.ID() != id {
		t.Errorf("expected %v, got %v", id, task.ID())
	}
	if task.Description() != description {
		t.Errorf("expected %v, got %v", description, task.Description())
	}
	if task.IsCompleted() {
		t.Errorf("expected false, got %v", task.IsCompleted())
	}
}

func TestDeleteTask(t *testing.T) {
	validUUID, _ := uuid.NewRandom()
	id, _ := task.NewTaskID(validUUID)
	description, _ := task.NewTaskDescription("Test task")
	task := task.NewTask(id, description)
	deletedTask := task.MarkAsComplete()

	if !deletedTask.IsCompleted() {
		t.Errorf("expected true, got %v", deletedTask.IsCompleted())
	}
	if deletedTask.UpdatedAt().Before(task.UpdatedAt()) {
		t.Errorf("expected updatedAt to be after %v, got %v", task.UpdatedAt(), deletedTask.UpdatedAt())
	}
}

func TestRestoreTask(t *testing.T) {
	validUUID, _ := uuid.NewRandom()
	id, _ := task.NewTaskID(validUUID)
	description, _ := task.NewTaskDescription("Test task")
	task := task.NewTask(id, description)
	deletedTask := task.MarkAsComplete()
	restoredTask := deletedTask.MarkAsIncomplete()

	if restoredTask.IsCompleted() {
		t.Errorf("expected false, got %v", restoredTask.IsCompleted())
	}
	if restoredTask.UpdatedAt().Before(deletedTask.UpdatedAt()) {
		t.Errorf("expected updatedAt to be after %v, got %v", deletedTask.UpdatedAt(), restoredTask.UpdatedAt())
	}
}

func TestChangeDescription(t *testing.T) {
	validUUID, _ := uuid.NewRandom()
	id, _ := task.NewTaskID(validUUID)
	description, _ := task.NewTaskDescription("Test task")
	newDescription, _ := task.NewTaskDescription("Updated task")
	task := task.NewTask(id, description)
	updatedTask := task.UpdateDescription(newDescription)

	if updatedTask.Description() != newDescription {
		t.Errorf("expected %v, got %v", newDescription, updatedTask.Description())
	}
	if updatedTask.UpdatedAt().Before(task.UpdatedAt()) {
		t.Errorf("expected updatedAt to be after %v, got %v", task.UpdatedAt(), updatedTask.UpdatedAt())
	}
}

func TestTaskTimestamps(t *testing.T) {
	validUUID, _ := uuid.NewRandom()
	id, _ := task.NewTaskID(validUUID)
	description, _ := task.NewTaskDescription("Test task")
	task := task.NewTask(id, description)

	if task.CreatedAt().IsZero() {
		t.Errorf("expected createdAt to be set, got %v", task.CreatedAt())
	}
	if task.UpdatedAt().IsZero() {
		t.Errorf("expected updatedAt to be set, got %v", task.UpdatedAt())
	}
	if !task.CreatedAt().Equal(task.UpdatedAt()) {
		t.Errorf("expected createdAt and updatedAt to be equal, got %v and %v", task.CreatedAt(), task.UpdatedAt())
	}
}
