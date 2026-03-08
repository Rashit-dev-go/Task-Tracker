package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Ошибка получения домашней директории: %v\n", err)
		os.Exit(1)
	}

	dataDir := filepath.Join(homeDir, ".task-tracker")
	dataFile := filepath.Join(dataDir, "tasks.json")

	app := &cli.Command{
		Name:  "task-tracker",
		Usage: "Простой трекер задач для личного использования",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Добавить новую задачу",
				Action: func(ctx *cli.Context) error {
					return addTask(ctx, dataFile)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "Показать список задач",
				Action: func(ctx *cli.Context) error {
					return listTasks(ctx, dataFile)
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "Отметить задачу как выполненную",
				Action: func(ctx *cli.Context) error {
					return completeTask(ctx, dataFile)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Удалить задачу",
				Action: func(ctx *cli.Context) error {
					return deleteTask(ctx, dataFile)
				},
			},
			{
				Name:  "init",
				Usage: "Инициализировать хранилище задач",
				Action: func(ctx *cli.Context) error {
					return initStorage(dataDir, dataFile)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}
}
