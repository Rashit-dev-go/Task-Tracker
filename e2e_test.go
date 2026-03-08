package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"task-tracker/testutils"
)

// P1-006: Complete task lifecycle (add→complete→delete)
func TestCompleteTaskLifecycle(t *testing.T) {
	// Arrange
	tempDir := testutils.SetupTempDir(t)
	homeDir := tempDir

	// Set environment variable for test
	os.Setenv("HOME", homeDir)
	defer os.Unsetenv("HOME")

	// Build the binary
	buildCmd := exec.Command("go", "build", "-o", filepath.Join(tempDir, "task-tracker"))
	buildCmd.Dir = tempDir
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	binaryPath := filepath.Join(tempDir, "task-tracker")

	// Initialize
	initCmd := exec.Command(binaryPath, "init")
	initCmd.Dir = tempDir
	err = initCmd.Run()
	if err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	// Add task (this would require interactive input simulation)
	// For now, we'll test the file operations directly
	t.Skip("E2E test requires interactive input simulation")

	// TODO: Add interactive input simulation for:
	// 1. Add task
	// 2. List tasks
	// 3. Complete task
	// 4. Delete task
}

// P2-008: Multiple tasks management workflow
func TestMultipleTasksManagement(t *testing.T) {
	// Arrange
	tempDir := testutils.SetupTempDir(t)
	homeDir := tempDir

	os.Setenv("HOME", homeDir)
	defer os.Unsetenv("HOME")

	// Create test data directly
	taskDir := filepath.Join(homeDir, ".task-tracker")
	err := os.MkdirAll(taskDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create task directory: %v", err)
	}

	// Create multiple tasks
	storage := TaskStorage{
		Tasks: []Task{
			{
				ID:          "task-1",
				Title:       "First Task",
				Description: "First task description",
				Status:      StatusTodo,
				CreatedAt:   time.Now().Add(-2 * time.Hour),
				UpdatedAt:   time.Now().Add(-2 * time.Hour),
			},
			{
				ID:          "task-2",
				Title:       "Second Task",
				Description: "Second task description",
				Status:      StatusInProgress,
				CreatedAt:   time.Now().Add(-1 * time.Hour),
				UpdatedAt:   time.Now().Add(-30 * time.Minute),
			},
			{
				ID:          "task-3",
				Title:       "Third Task",
				Description: "Third task description",
				Status:      StatusDone,
				CreatedAt:   time.Now().Add(-3 * time.Hour),
				UpdatedAt:   time.Now().Add(-1 * time.Hour),
				CompletedAt: &[]time.Time{time.Now().Add(-1 * time.Hour)}[0],
			},
		},
	}

	dataFile := filepath.Join(taskDir, "tasks.json")
	err = saveTasks(&storage, dataFile)
	if err != nil {
		t.Fatalf("Failed to save test tasks: %v", err)
	}

	// Act - Load tasks
	loadedStorage, err := loadTasks(dataFile)
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	// Assert
	if len(loadedStorage.Tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(loadedStorage.Tasks))
	}

	// Check task statuses
	todoCount := 0
	inProgressCount := 0
	doneCount := 0

	for _, task := range loadedStorage.Tasks {
		switch task.Status {
		case StatusTodo:
			todoCount++
		case StatusInProgress:
			inProgressCount++
		case StatusDone:
			doneCount++
		}
	}

	if todoCount != 1 {
		t.Errorf("Expected 1 todo task, got %d", todoCount)
	}
	if inProgressCount != 1 {
		t.Errorf("Expected 1 in-progress task, got %d", inProgressCount)
	}
	if doneCount != 1 {
		t.Errorf("Expected 1 done task, got %d", doneCount)
	}
}

// P2-003: Error recovery workflow
func TestErrorRecoveryWorkflow(t *testing.T) {
	// Arrange
	tempDir := testutils.SetupTempDir(t)
	dataFile := filepath.Join(tempDir, "corrupted.json")

	// Create corrupted JSON file
	corruptedContent := []byte(`{"tasks": [{"id": "1", "title": "Task 1", "status": "todo"`)
	err := os.WriteFile(dataFile, corruptedContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create corrupted file: %v", err)
	}

	// Act - Try to load corrupted data
	storage, err := loadTasks(dataFile)

	// Assert
	if err == nil {
		t.Error("Expected error when loading corrupted JSON, but got none")
	}
	if storage != nil {
		t.Error("Expected nil storage when error occurs")
	}
}

// P3-002: Error recovery workflow (extended)
func TestExtendedErrorRecovery(t *testing.T) {
	// Test various error scenarios
	testCases := []struct {
		name        string
		createFile  func(string) error
		expectError bool
	}{
		{
			name: "Empty file",
			createFile: func(path string) error {
				return os.WriteFile(path, []byte(""), 0644)
			},
			expectError: true,
		},
		{
			name: "Invalid JSON structure",
			createFile: func(path string) error {
				return os.WriteFile(path, []byte(`{"invalid": "structure"}`), 0644)
			},
			expectError: false, // Should handle gracefully
		},
		{
			name: "Valid empty tasks",
			createFile: func(path string) error {
				return os.WriteFile(path, []byte(`{"tasks": []}`), 0644)
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := testutils.SetupTempDir(t)
			dataFile := filepath.Join(tempDir, "test.json")

			err := tc.createFile(dataFile)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			storage, err := loadTasks(dataFile)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for %s, but got none", tc.name)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Expected no error for %s, got %v", tc.name, err)
			}
			if storage == nil && !tc.expectError {
				t.Error("Expected non-nil storage when no error occurs")
			}
		})
	}
}
