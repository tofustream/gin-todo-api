package task

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	id, _ := NewTaskID("1")
	description, _ := NewTaskDescription("Test task")
	task := NewTask(id, description)

	if task.ID() != id {
		t.Errorf("expected %v, got %v", id, task.ID())
	}
	if task.Description() != description {
		t.Errorf("expected %v, got %v", description, task.Description())
	}
	if task.IsDeleted() {
		t.Errorf("expected false, got %v", task.IsDeleted())
	}
}

func TestDeleteTask(t *testing.T) {
	id, _ := NewTaskID("1")
	description, _ := NewTaskDescription("Test task")
	task := NewTask(id, description)
	deletedTask := task.Delete()

	if !deletedTask.IsDeleted() {
		t.Errorf("expected true, got %v", deletedTask.IsDeleted())
	}
	if deletedTask.UpdatedAt().Before(task.UpdatedAt()) {
		t.Errorf("expected updatedAt to be after %v, got %v", task.UpdatedAt(), deletedTask.UpdatedAt())
	}
}

func TestRestoreTask(t *testing.T) {
	id, _ := NewTaskID("1")
	description, _ := NewTaskDescription("Test task")
	task := NewTask(id, description)
	deletedTask := task.Delete()
	restoredTask := deletedTask.Restore()

	if restoredTask.IsDeleted() {
		t.Errorf("expected false, got %v", restoredTask.IsDeleted())
	}
	if restoredTask.UpdatedAt().Before(deletedTask.UpdatedAt()) {
		t.Errorf("expected updatedAt to be after %v, got %v", deletedTask.UpdatedAt(), restoredTask.UpdatedAt())
	}
}

func TestChangeDescription(t *testing.T) {
	id, _ := NewTaskID("1")
	description, _ := NewTaskDescription("Test task")
	newDescription, _ := NewTaskDescription("Updated task")
	task := NewTask(id, description)
	updatedTask := task.ChangeDescription(newDescription)

	if updatedTask.Description() != newDescription {
		t.Errorf("expected %v, got %v", newDescription, updatedTask.Description())
	}
	if updatedTask.UpdatedAt().Before(task.UpdatedAt()) {
		t.Errorf("expected updatedAt to be after %v, got %v", task.UpdatedAt(), updatedTask.UpdatedAt())
	}
}

func TestTaskTimestamps(t *testing.T) {
	id, _ := NewTaskID("1")
	description, _ := NewTaskDescription("Test task")
	task := NewTask(id, description)

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
