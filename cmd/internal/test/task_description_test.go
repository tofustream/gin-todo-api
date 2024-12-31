package task_test

import (
	"testing"

	"github.com/tofustream/gin-todo-api/cmd/internal/task"
)

func TestNewTaskDescription(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"Empty description", "", true},
		{"Valid description", "This is a valid task description.", false},
		{"Too long description", "This description is way too long and exceeds the maximum allowed length of 140 characters!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := task.NewTaskDescription(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTaskDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskDescription_Value(t *testing.T) {
	desc, err := task.NewTaskDescription("This is a valid task description.")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if got := desc.Value(); got != "This is a valid task description." {
		t.Errorf("TaskDescription.Value() = %v, want %v", got, "This is a valid task description.")
	}
}
