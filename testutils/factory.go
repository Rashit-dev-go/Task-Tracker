package testutils

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

// Task и Status типы определены в main пакете
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type TaskStorage struct {
	Tasks []Task `json:"tasks"`
}

// Status константы
const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

// TaskFactory создает тестовые задачи с возможностью переопределения полей
func CreateTestTask(t *testing.T, overrides ...func(*Task)) Task {
	task := Task{
		ID:          uuid.New().String(),
		Title:       "Test Task " + uuid.New().String()[:8],
		Description: "Test description",
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	for _, override := range overrides {
		override(&task)
	}

	return task
}

// CreateTestTaskWithSpecificTitle создает задачу с указанным заголовком
func CreateTestTaskWithSpecificTitle(t *testing.T, title string) Task {
	return CreateTestTask(t, func(task *Task) {
		task.Title = title
	})
}

// CreateTestTaskWithStatus создает задачу с указанным статусом
func CreateTestTaskWithStatus(t *testing.T, status string) Task {
	return CreateTestTask(t, func(task *Task) {
		task.Status = status
	})
}

// CreateMultipleTestTasks создает несколько тестовых задач
func CreateMultipleTestTasks(t *testing.T, count int) []Task {
	tasks := make([]Task, count)
	for i := 0; i < count; i++ {
		tasks[i] = CreateTestTask(t, func(task *Task) {
			task.Title = "Test Task " + string(rune(i+1))
		})
	}
	return tasks
}

// CreateTestTaskStorage создает тестовое хранилище с задачами
func CreateTestTaskStorage(t *testing.T, taskCount int) TaskStorage {
	tasks := CreateMultipleTestTasks(t, taskCount)
	return TaskStorage{
		Tasks: tasks,
	}
}

// CreateEmptyTaskStorage создает пустое тестовое хранилище
func CreateEmptyTaskStorage(t *testing.T) TaskStorage {
	return TaskStorage{
		Tasks: []Task{},
	}
}
