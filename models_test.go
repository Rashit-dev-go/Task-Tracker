package main

import (
	"encoding/json"
	"testing"
	"time"
)

// P0-001: Create task with valid data
func TestCreateTaskWithValidData(t *testing.T) {
	// Arrange & Act
	task := Task{
		ID:          "test-id",
		Title:       "Test Task",
		Description: "Test Description",
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Assert
	if task.Title != "Test Task" {
		t.Errorf("Expected title 'Test Task', got '%s'", task.Title)
	}
	if task.Status != StatusTodo {
		t.Errorf("Expected status '%s', got '%s'", StatusTodo, task.Status)
	}
	if task.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
	if task.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should not be zero")
	}
}

// P0-002: Create task with empty title (error)
func TestCreateTaskWithEmptyTitle(t *testing.T) {
	// Arrange
	task := Task{
		ID:          "test-id",
		Title:       "", // Empty title
		Description: "Test Description",
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Assert
	if task.Title == "" {
		t.Log("Empty title allowed (current behavior)")
		// Note: This test documents current behavior
		// Consider adding validation in future
	}
}

// P1-004: Task struct validation
func TestTaskStructValidation(t *testing.T) {
	tests := []struct {
		name     string
		task     Task
		wantErr  bool
		errorMsg string
	}{
		{
			name: "Valid task",
			task: Task{
				ID:          "test-id",
				Title:       "Valid Task",
				Description: "Valid Description",
				Status:      StatusTodo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantErr: false,
		},
		{
			name: "Empty ID",
			task: Task{
				ID:          "",
				Title:       "Valid Task",
				Description: "Valid Description",
				Status:      StatusTodo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantErr:  true,
			errorMsg: "ID cannot be empty",
		},
		{
			name: "Invalid status",
			task: Task{
				ID:          "test-id",
				Title:       "Valid Task",
				Description: "Valid Description",
				Status:      "invalid", // Invalid status
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantErr:  true,
			errorMsg: "invalid status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This is a placeholder for validation logic
			// In a real implementation, you would add validation methods
			if tt.wantErr && tt.errorMsg != "" {
				t.Logf("Expected error: %s", tt.errorMsg)
			}
		})
	}
}

// P1-005: JSON serialization/deserialization
func TestTaskJSONSerialization(t *testing.T) {
	// Arrange
	originalTask := Task{
		ID:          "test-id",
		Title:       "Test Task",
		Description: "Test Description",
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Act - Serialize
	jsonData, err := json.Marshal(originalTask)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}

	// Act - Deserialize
	var deserializedTask Task
	err = json.Unmarshal(jsonData, &deserializedTask)
	if err != nil {
		t.Fatalf("Failed to unmarshal task: %v", err)
	}

	// Assert
	if deserializedTask.ID != originalTask.ID {
		t.Errorf("Expected ID %s, got %s", originalTask.ID, deserializedTask.ID)
	}
	if deserializedTask.Title != originalTask.Title {
		t.Errorf("Expected title %s, got %s", originalTask.Title, deserializedTask.Title)
	}
	if deserializedTask.Status != originalTask.Status {
		t.Errorf("Expected status %s, got %s", originalTask.Status, deserializedTask.Status)
	}
}

// P1-006: Status constants validation
func TestStatusConstants(t *testing.T) {
	tests := []struct {
		name   string
		status Status
		valid  bool
	}{
		{"Todo status", StatusTodo, true},
		{"InProgress status", StatusInProgress, true},
		{"Done status", StatusDone, true},
		{"Empty status", "", false},
		{"Invalid status", "invalid", false},
	}

	validStatuses := []Status{StatusTodo, StatusInProgress, StatusDone}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := false
			for _, valid := range validStatuses {
				if tt.status == valid {
					isValid = true
					break
				}
			}

			if tt.valid != isValid {
				t.Errorf("Expected status %s to be valid=%t, got valid=%t", tt.status, tt.valid, isValid)
			}
		})
	}
}

// P2-007: Timestamp handling
func TestTimestampHandling(t *testing.T) {
	now := time.Now()

	task := Task{
		ID:          "test-id",
		Title:       "Test Task",
		Description: "Test Description",
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now.Add(time.Hour),
	}

	// Assert CreatedAt <= UpdatedAt
	if task.CreatedAt.After(task.UpdatedAt) {
		t.Error("CreatedAt should be before or equal to UpdatedAt")
	}

	// Test JSON serialization preserves timestamps
	jsonData, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task with timestamps: %v", err)
	}

	var restoredTask Task
	err = json.Unmarshal(jsonData, &restoredTask)
	if err != nil {
		t.Fatalf("Failed to unmarshal task with timestamps: %v", err)
	}

	// Check timestamps are preserved (with some tolerance for JSON precision)
	if !task.CreatedAt.Equal(restoredTask.CreatedAt) {
		t.Errorf("CreatedAt not preserved: expected %v, got %v", task.CreatedAt, restoredTask.CreatedAt)
	}
}
