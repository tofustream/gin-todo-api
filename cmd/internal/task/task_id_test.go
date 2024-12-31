package task

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewTaskID(t *testing.T) {
	validID, _ := uuid.NewRandom()
	invalidID := uuid.Nil
	tests := []struct {
		value    uuid.UUID
		expected uuid.UUID
		hasError bool
	}{
		{value: validID, expected: validID, hasError: false},
		{value: invalidID, expected: invalidID, hasError: true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			taskID, err := NewTaskID(tt.value)
			if tt.hasError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got %v", err)
				}
				if taskID.Value() != tt.expected && tt.value != invalidID {
					t.Errorf("expected %v but got %v", tt.expected, taskID.Value())
				}
			}
		})
	}
}

func TestTaskID_Value(t *testing.T) {
	validID, _ := uuid.NewRandom()
	taskID, err := NewTaskID(validID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if taskID.Value() != validID {
		t.Errorf("expected %v but got %v", validID, taskID.Value())
	}
}
