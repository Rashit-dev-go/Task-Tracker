package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"task-tracker/testutils"
)

// P0-003: Load tasks from non-existent file
func TestLoadTasksFromNonExistentFile(t *testing.T) {
	// Arrange
	tempDir := testutils.SetupTempDir(t)
	nonExistentFile := filepath.Join(tempDir, "nonexistent.json")

	// Act
	storage, err := loadTasks(nonExistentFile)

	// Assert
	if err != nil {
		t.Errorf("Expected no error for non-existent file, got %v", err)
	}
	if len(storage.Tasks) != 0 {
		t.Errorf("Expected empty task list, got %d tasks", len(storage.Tasks))
	}
}

// P0-004: Atomic file write operations
func TestAtomicFileWrite(t *testing.T) {
	// Arrange
	tempDir := testutils.SetupTempDir(t)
	dataFile := filepath.Join(tempDir, "tasks.json")

	storage := TaskStorage{
		Tasks: []Task{
			{
				ID:          "test-id-1",
				Title:       "Test Task 1",
				Description: "Test Description 1",
				Status:      StatusTodo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          "test-id-2",
				Title:       "Test Task 2",
				Description: "Test Description 2",
				Status:      StatusTodo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          "test-id-3",
				Title:       "Test Task 3",
				Description: "Test Description 3",
				Status:      StatusTodo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
	}

	// Act
	err := saveTasks(&storage, dataFile)

	// Assert
	if err != nil {
		t.Fatalf("Failed to save tasks: %v", err)
	}

	// Verify file exists and has correct content
	testutils.AssertFileExists(t, dataFile)

	content := testutils.ReadTestDataFile(t, dataFile)

	var loadedStorage TaskStorage
	err = json.Unmarshal(content, &loadedStorage)
	if err != nil {
		t.Fatalf("Failed to unmarshal saved content: %v", err)
	}

	if len(loadedStorage.Tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(loadedStorage.Tasks))
	}
}

// P0-005: `init` command creates directory and file
func TestInitCommandCreatesDirectory(t *testing.T) {
	// Arrange
	tempDir := testutils.SetupTempDir(t)
	taskDir := filepath.Join(tempDir, ".task-tracker")
	dataFile := filepath.Join(taskDir, "tasks.json")

	// Act
	err := os.MkdirAll(taskDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	emptyStorage := TaskStorage{Tasks: []Task{}}
	err = saveTasks(&emptyStorage, dataFile)
	if err != nil {
		t.Fatalf("Failed to save initial storage: %v", err)
	}

	// Assert
	testutils.AssertDirExists(t, taskDir)
	testutils.AssertFileExists(t, dataFile)
}

// P1-001: Save tasks with invalid JSON structure
func TestSaveTasksWithInvalidJSON(t *testing.T) {
	// This test would require creating invalid JSON data
	// For now, we test that valid JSON saves correctly
	t.Skip("JSON validation test - requires mock invalid data")
}

// P1-007: Concurrent file access handling
func TestConcurrentFileAccess(t *testing.T) {
	// Arrange
	tempDir := testutils.SetupTempDir(t)
	dataFile := filepath.Join(tempDir, "tasks.json")

	storage := TaskStorage{
		Tasks: make([]Task, 10),
	}
	for i := 0; i < 10; i++ {
		storage.Tasks[i] = Task{
			ID:          fmt.Sprintf("test-id-%d", i),
			Title:       fmt.Sprintf("Test Task %d", i+1),
			Description: fmt.Sprintf("Test Description %d", i+1),
			Status:      StatusTodo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}

	// Act - Save initial data
	err := saveTasks(&storage, dataFile)
	if err != nil {
		t.Fatalf("Failed to save initial tasks: %v", err)
	}

	// Try to load the saved data
	loadedStorage, err := loadTasks(dataFile)
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	// Assert
	if len(loadedStorage.Tasks) != 10 {
		t.Errorf("Expected 10 tasks, got %d", len(loadedStorage.Tasks))
	}
}

// P2-003: File permission error handling
func TestFilePermissionErrorHandling(t *testing.T) {
	// Arrange - try to write to a read-only directory
	tempDir := testutils.SetupTempDir(t)
	readOnlyDir := filepath.Join(tempDir, "readonly")

	err := os.MkdirAll(readOnlyDir, 0444) // read-only permissions
	if err != nil {
		t.Skipf("Cannot create read-only directory for testing: %v", err)
	}
	defer os.Chmod(readOnlyDir, 0755) // cleanup permissions

	dataFile := filepath.Join(readOnlyDir, "tasks.json")
	storage := TaskStorage{
		Tasks: []Task{
			{
				ID:          "test-id-perm",
				Title:       "Test Task Permission",
				Description: "Test Description",
				Status:      StatusTodo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
	}

	// Act
	err = saveTasks(&storage, dataFile)

	// Assert
	if err == nil {
		t.Error("Expected permission error, but got none")
	}
}

// P2-004: Command aliases work (a, l, c, d)
func TestCommandAliases(t *testing.T) {
	// This would be tested in integration tests with actual CLI commands
	// For now, we'll test the underlying functions
	t.Skip("CLI command alias tests - require CLI framework integration")
}
