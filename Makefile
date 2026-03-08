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
	@echo "🧪 Запуск всех тестов..."
	@go test -v ./...

# Тестирование с покрытием
test-coverage:
	@echo "📊 Запуск тестов с покрытием..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "📈 Отчет покрытия: coverage.html"

# Только P0 тесты (критические)
test-p0:
	@echo "🔥 Запуск P0 тестов..."
	@go test -v -run "P0" ./...

# P0 + P1 тесты
test-core:
	@echo "⚡ Запуск P0+P1 тестов..."
	@go test -v -run "(P0|P1)" ./...

# Performance тесты
test-bench:
	@echo "🚀 Запуск performance тестов..."
	@go test -v -bench=. ./...

# Тесты с детальной информацией
test-verbose:
	@echo "🔍 Детальный запуск тестов..."
	@go test -v -race -cover ./...

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
	@echo "  make test         - Запуск всех тестов"
	@echo "  make test-coverage- Тесты с покрытием кода"
	@echo "  make test-p0      - Только P0 (критические) тесты"
	@echo "  make test-core    - P0+P1 тесты"
	@echo "  make test-bench   - Performance тесты"
	@echo "  make test-verbose - Детальный запуск тестов"
	@echo "  make fmt      - Форматирование кода"
	@echo "  make vet      - Проверка кода"
	@echo "  make init     - Инициализация хранилища"
	@echo "  make demo     - Демонстрация работы"
	@echo "  make check    - Полная проверка"
	@echo "  make help     - Эта справка"
