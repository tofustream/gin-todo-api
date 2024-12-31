package task

import (
	"testing"
)

func TestNewTaskID(t *testing.T) {
	tests := []struct {
		value    string
		expected string
		hasError bool
	}{
		{value: "1", expected: "1", hasError: false},
		{value: "", expected: "", hasError: true},
		{value: "100", expected: "100", hasError: false},
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
				if taskID.Value() != tt.expected && tt.value != "" {
					t.Errorf("expected %v but got %v", tt.expected, taskID.Value())
				}
			}
		})
	}
}

func TestTaskID_Value(t *testing.T) {
	taskID, err := NewTaskID("10")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if taskID.Value() != "10" {
		t.Errorf("expected 10 but got %v", taskID.Value())
	}
}
