# Task Tracker Makefile

# Переменные
BINARY_NAME=task-tracker
BUILD_DIR=build
GO_FILES=$(shell find . -name "*.go" -type f)

# Основные команды
.PHONY: all build clean run install test help

all: build

# Сборка приложения
build:
	@echo "🔨 Сборка $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_FILES)
	@echo "✅ Сборка завершена: $(BUILD_DIR)/$(BINARY_NAME)"

# Запуск приложения
run: build
	@echo "🚀 Запуск $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME) --help

# Установка в систему
install: build
	@echo "📦 Установка $(BINARY_NAME) в /usr/local/bin..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "✅ Установлено: /usr/local/bin/$(BINARY_NAME)"

# Очистка
clean:
	@echo "🧹 Очистка..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "✅ Очистка завершена"

# Тестирование
test:
	@echo "🧪 Запуск тестов..."
	@go test -v ./...

# Форматирование кода
fmt:
	@echo "📝 Форматирование кода..."
	@go fmt ./...

# Проверка кода
vet:
	@echo "🔍 Проверка кода..."
	@go vet ./...

# Инициализация проекта
init:
	@echo "🔧 Инициализация хранилища..."
	@go run $(GO_FILES) init

# Добавление тестовой задачи
demo:
	@echo "📋 Демонстрация: добавление тестовых задач..."
	@go run $(GO_FILES) add
	@go run $(GO_FILES) list

# Полная проверка
check: fmt vet test
	@echo "✅ Все проверки пройдены"

# Справка
help:
	@echo "Доступные команды:"
	@echo "  make build    - Сборка приложения"
	@echo "  make run      - Запуск приложения"
	@echo "  make install  - Установка в систему"
	@echo "  make clean    - Очистка"
	@echo "  make test     - Запуск тестов"
	@echo "  make fmt      - Форматирование кода"
	@echo "  make vet      - Проверка кода"
	@echo "  make init     - Инициализация хранилища"
	@echo "  make demo     - Демонстрация работы"
	@echo "  make check    - Полная проверка"
	@echo "  make help     - Эта справка"
