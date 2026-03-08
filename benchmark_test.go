package main

import (
	"fmt"
	"testing"
	"time"

	"task-tracker/testutils"
)

// P2-001: Load 1000+ tasks performance
func BenchmarkLoad1000Tasks(b *testing.B) {
	tempDir := testutils.SetupTempDir(&testing.T{})
	dataFile := fmt.Sprintf("%s/large_tasks.json", tempDir)

	// Create 1000 tasks
	storage := TaskStorage{
		Tasks: make([]Task, 1000),
	}

	for i := 0; i < 1000; i++ {
		storage.Tasks[i] = Task{
			ID:          fmt.Sprintf("task-%d", i),
			Title:       fmt.Sprintf("Task %d", i+1),
			Description: fmt.Sprintf("Description for task %d", i+1),
			Status:      StatusTodo,
			CreatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
			UpdatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
		}
	}

	err := saveTasks(&storage, dataFile)
	if err != nil {
		b.Fatalf("Failed to save test data: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := loadTasks(dataFile)
		if err != nil {
			b.Fatalf("Failed to load tasks: %v", err)
		}
	}
}

// P2-002: Large task list display performance
func BenchmarkTaskListDisplay(b *testing.B) {
	// Create large task list
	storage := TaskStorage{
		Tasks: make([]Task, 1000),
	}

	for i := 0; i < 1000; i++ {
		storage.Tasks[i] = Task{
			ID:          fmt.Sprintf("task-%d", i),
			Title:       fmt.Sprintf("Task %d", i+1),
			Description: fmt.Sprintf("Description for task %d", i+1),
			Status:      StatusTodo,
			CreatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
			UpdatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Simulate task list processing
		taskCount := len(storage.Tasks)
		todoCount := 0
		inProgressCount := 0
		doneCount := 0

		for _, task := range storage.Tasks {
			switch task.Status {
			case StatusTodo:
				todoCount++
			case StatusInProgress:
				inProgressCount++
			case StatusDone:
				doneCount++
			}
		}

		// Prevent optimization
		if taskCount != 1000 || todoCount+inProgressCount+doneCount != 1000 {
			b.Error("Unexpected task counts")
		}
	}
}

// P2-005: Save performance with large datasets
func BenchmarkSave1000Tasks(b *testing.B) {
	// Create 1000 tasks
	storage := TaskStorage{
		Tasks: make([]Task, 1000),
	}

	for i := 0; i < 1000; i++ {
		storage.Tasks[i] = Task{
			ID:          fmt.Sprintf("task-%d", i),
			Title:       fmt.Sprintf("Task %d", i+1),
			Description: fmt.Sprintf("Description for task %d", i+1),
			Status:      StatusTodo,
			CreatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
			UpdatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tempDir := testutils.SetupTempDir(&testing.T{})
		testFile := fmt.Sprintf("%s/benchmark_save_%d.json", tempDir, i)
		err := saveTasks(&storage, testFile)
		if err != nil {
			b.Fatalf("Failed to save tasks: %v", err)
		}
	}
}

// P3-003: Memory usage with 10k tasks
func BenchmarkMemoryUsage10KTasks(b *testing.B) {
	// Create 10,000 tasks
	storage := TaskStorage{
		Tasks: make([]Task, 10000),
	}

	for i := 0; i < 10000; i++ {
		storage.Tasks[i] = Task{
			ID:          fmt.Sprintf("task-%d", i),
			Title:       fmt.Sprintf("Task %d", i+1),
			Description: fmt.Sprintf("Description for task %d", i+1),
			Status:      StatusTodo,
			CreatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
			UpdatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Simulate memory-intensive operations
		taskCount := len(storage.Tasks)

		// Process all tasks
		for _, task := range storage.Tasks {
			_ = task.ID
			_ = task.Title
			_ = task.Status
			_ = task.CreatedAt
			_ = task.UpdatedAt
		}

		// Prevent optimization
		if taskCount != 10000 {
			b.Error("Unexpected task count")
		}
	}
}

// P3-001: File size impact on operations
func BenchmarkFileSizeImpact(b *testing.B) {
	sizes := []int{100, 500, 1000, 5000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			tempDir := testutils.SetupTempDir(&testing.T{})
			dataFile := fmt.Sprintf("%s/size_test_%d.json", tempDir, size)

			// Create tasks
			storage := TaskStorage{
				Tasks: make([]Task, size),
			}

			for i := 0; i < size; i++ {
				storage.Tasks[i] = Task{
					ID:          fmt.Sprintf("task-%d", i),
					Title:       fmt.Sprintf("Task %d", i+1),
					Description: fmt.Sprintf("Description for task %d", i+1),
					Status:      StatusTodo,
					CreatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
					UpdatedAt:   time.Now().Add(time.Duration(-i) * time.Minute),
				}
			}

			err := saveTasks(&storage, dataFile)
			if err != nil {
				b.Fatalf("Failed to save test data: %v", err)
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, err := loadTasks(dataFile)
				if err != nil {
					b.Fatalf("Failed to load tasks: %v", err)
				}
			}
		})
	}
}
