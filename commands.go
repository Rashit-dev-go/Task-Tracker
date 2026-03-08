package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func addTask(ctx *cli.Context, dataFile string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Название задачи: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	if title == "" {
		return fmt.Errorf("название задачи не может быть пустым")
	}

	fmt.Print("Описание (опционально): ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	storage, err := loadTasks(dataFile)
	if err != nil {
		return err
	}

	task := createTask(title, description)
	storage.Tasks = append(storage.Tasks, task)

	if err := saveTasks(storage, dataFile); err != nil {
		return err
	}

	fmt.Printf("✅ Задача добавлена: %s (ID: %s)\n", task.Title, task.ID)
	return nil
}

func listTasks(ctx *cli.Context, dataFile string) error {
	storage, err := loadTasks(dataFile)
	if err != nil {
		return err
	}

	if len(storage.Tasks) == 0 {
		fmt.Println("📝 Задач пока нет")
		return nil
	}

	fmt.Printf("\n📋 Список задач (%d):\n\n", len(storage.Tasks))

	for _, task := range storage.Tasks {
		statusIcon := getStatusIcon(task.Status)
		fmt.Printf("%s [%s] %s\n", statusIcon, task.ID[:8], task.Title)

		if task.Description != "" {
			fmt.Printf("   📄 %s\n", task.Description)
		}

		fmt.Printf("   📅 Создана: %s", task.CreatedAt.Format("02.01.2006 15:04"))

		if task.Status == StatusDone && task.CompletedAt != nil {
			fmt.Printf(" | Выполнена: %s", task.CompletedAt.Format("02.01.2006 15:04"))
		}

		fmt.Println("\n")
	}

	return nil
}

func completeTask(ctx *cli.Context, dataFile string) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("укажите ID задачи для отметки как выполненной")
	}

	taskID := ctx.Args().First()

	storage, err := loadTasks(dataFile)
	if err != nil {
		return err
	}

	task := findTaskByID(storage.Tasks, taskID)
	if task == nil {
		return fmt.Errorf("задача с ID %s не найдена", taskID)
	}

	if task.Status == StatusDone {
		fmt.Printf("ℹ️  Задача уже выполнена: %s\n", task.Title)
		return nil
	}

	task.Status = StatusDone
	task.UpdatedAt = time.Now()
	now := time.Now()
	task.CompletedAt = &now

	if err := saveTasks(storage, dataFile); err != nil {
		return err
	}

	fmt.Printf("✅ Задача выполнена: %s\n", task.Title)
	return nil
}

func deleteTask(ctx *cli.Context, dataFile string) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("укажите ID задачи для удаления")
	}

	taskID := ctx.Args().First()

	storage, err := loadTasks(dataFile)
	if err != nil {
		return err
	}

	taskIndex := -1
	for i, task := range storage.Tasks {
		if task.ID == taskID {
			taskIndex = i
			break
		}
	}

	if taskIndex == -1 {
		return fmt.Errorf("задача с ID %s не найдена", taskID)
	}

	taskTitle := storage.Tasks[taskIndex].Title
	storage.Tasks = append(storage.Tasks[:taskIndex], storage.Tasks[taskIndex+1:]...)

	if err := saveTasks(storage, dataFile); err != nil {
		return err
	}

	fmt.Printf("🗑️  Задача удалена: %s\n", taskTitle)
	return nil
}

func getStatusIcon(status Status) string {
	switch status {
	case StatusTodo:
		return "⭕"
	case StatusInProgress:
		return "🔄"
	case StatusDone:
		return "✅"
	default:
		return "❓"
	}
}
