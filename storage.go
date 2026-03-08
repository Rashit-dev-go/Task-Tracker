package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func initStorage(dataDir, dataFile string) error {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("ошибка создания директории %s: %w", dataDir, err)
	}

	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		storage := TaskStorage{
			Tasks: []Task{},
		}
		
		data, err := json.MarshalIndent(storage, "", "  ")
		if err != nil {
			return fmt.Errorf("ошибка кодирования данных: %w", err)
		}
		
		if err := os.WriteFile(dataFile, data, 0644); err != nil {
			return fmt.Errorf("ошибка создания файла %s: %w", dataFile, err)
		}
		
		fmt.Printf("✅ Хранилище инициализировано: %s\n", dataFile)
	} else {
		fmt.Printf("ℹ️  Хранилище уже существует: %s\n", dataFile)
	}
	
	return nil
}

func loadTasks(dataFile string) (*TaskStorage, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return &TaskStorage{Tasks: []Task{}}, nil
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %w", dataFile, err)
	}

	var storage TaskStorage
	if err := json.Unmarshal(data, &storage); err != nil {
		return nil, fmt.Errorf("ошибка декодирования данных: %w", err)
	}

	return &storage, nil
}

func saveTasks(storage *TaskStorage, dataFile string) error {
	if err := os.MkdirAll(filepath.Dir(dataFile), 0755); err != nil {
		return fmt.Errorf("ошибка создания директории: %w", err)
	}

	data, err := json.MarshalIndent(storage, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка кодирования данных: %w", err)
	}

	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		return fmt.Errorf("ошибка записи файла %s: %w", dataFile, err)
	}

	return nil
}

func createTask(title, description string) Task {
	now := time.Now()
	return Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func findTaskByID(tasks []Task, id string) *Task {
	for i, task := range tasks {
		if task.ID == id {
			return &tasks[i]
		}
	}
	return nil
}
