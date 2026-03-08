# Task Tracker Test Framework

**Framework:** Go Test (встроенный)  
**Project Type:** Backend CLI приложение  
**Test Stack:** Unit + Integration + E2E + Benchmark

---

## 🚀 Быстрый старт

### Запуск всех тестов
```bash
make test
```

### Запуск с покрытием кода
```bash
make test-coverage
# Отчет откроется в coverage.html
```

### Запуск только критических тестов
```bash
make test-p0        # Только P0 (критические)
make test-core      # P0 + P1 (основные)
make test-bench     # Performance тесты
```

---

## 📋 Архитектура тестов

### Структура директорий
```
├── models_test.go          # Unit тесты для моделей
├── storage_test.go         # Integration тесты хранилища  
├── e2e_test.go            # End-to-end тесты
├── benchmark_test.go      # Performance тесты
├── testutils/
│   ├── factory.go         # Фабрики тестовых данных
│   └── helpers.go         # Вспомогательные функции
└── testdata/
    └── sample_tasks.json  # Пример тестовых данных
```

### Типы тестов

#### **Unit тесты (P0-P1)**
- Валидация моделей данных
- JSON сериализация/десериализация
- Константы и статусы

#### **Integration тесты (P0-P2)**
- Файловые операции
- Атомарная запись/чтение
- Обработка ошибок файловой системы

#### **E2E тесты (P1-P3)**
- Полный жизненный цикл задач
- Управление множественными задачами
- Восстановление после ошибок

#### **Benchmark тесты (P2-P3)**
- Производительность с большими наборами данных
- Использование памяти
- Влияние размера файла на операции

---

## 🏗️ Фабрики и хелперы

### Task Factory
```go
// Создание тестовой задачи
task := testutils.CreateTestTask(t)

// С определенным заголовком
task := testutils.CreateTestTaskWithSpecificTitle(t, "My Task")

// С определенным статусом
task := testutils.CreateTestTaskWithStatus(t, StatusDone)

// Несколько задач
tasks := testutils.CreateMultipleTestTasks(t, 10)
```

### File Helpers
```go
// Настройка временной директории
tempDir := testutils.SetupTempDir(t)

// Проверка существования файла
testutils.AssertFileExists(t, filePath)

// Создание тестовых данных
testutils.CreateTestDataFile(t, filePath, content)
```

---

## 🎯 Приоритеты тестов

### **P0 (Критические)** - 8 тестов
- Блокируют core функциональность
- Высокий риск (score ≥6)
- Нет обходных путей

**Примеры:**
- Создание задачи с валидными данными
- Атомарные файловые операции
- Базовые CLI команды

### **P1 (Высокие)** - 7 тестов  
- Важные функции
- Средний риск (score 3-4)
- Сложные обходные пути

**Примеры:**
- Валидация JSON структуры
- Обработка ошибок
- Полный жизненный цикл задачи

### **P2 (Средние)** - 8 тестов
- Второстепенные функции
- Низкий риск (score 1-2)
- Edge cases

**Примеры:**
- Performance с 1000+ задач
- Обработка разрешений файлов
- Cross-platform совместимость

### **P3 (Низкие)** - 3 теста
- Улучшения
- Exploratory тесты
- Бенчмарки

**Примеры:**
- Memory usage с 10k задач
- Восстановление после ошибок
- Исследовательское тестирование

---

## 🔧 Локальная разработка

### Environment Setup
```bash
# Копирование конфигурации
cp .env.example .env

# Настройка переменных
export TEST_ENV=test
export TASK_TRACKER_DATA_DIR=./testdata/temp
```

### Debug режим
```bash
# Детальный вывод тестов
make test-verbose

# С race condition детектором
go test -race -v ./...
```

### Изолированное тестирование
```bash
# Тесты используют автоматические временные директории
# Никаких конфликтов с локальными данными
# Автоматическая cleanup после каждого теста
```

---

## 🚀 CI Integration

### GitHub Actions (рекомендация)
```yaml
name: Tests
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: '1.21'
      - name: Run tests
        run: make test-core
      - name: Run benchmarks  
        run: make test-bench
```

### Quality Gates
- ✅ P0 тесты: 100% pass rate (blocking)
- ✅ P1 тесты: ≥95% pass rate  
- ✅ Coverage: ≥80%
- ✅ Performance: 1000 задач < 2с

---

## 📚 Best Practices

### ✅ Рекомендации
- Используйте `t.TempDir()` для изоляции
- Применяйте factory pattern для тестовых данных
- Следуйте Given/When/Then структуре
- Добавляйте descriptive test names

### ❌ Избегайте
- Hardcoded путей к файлам
- Зависимости от порядка выполнения тестов
- Shared state между тестами
- Слишком сложных тестовых сценариев

---

## 🔍 Отладка

### Поиск проблем
```bash
# Детальный вывод конкретного теста
go test -v -run TestSpecificFunction

# С coverage для конкретного пакета
go test -cover -v ./...

# Performance профилирование
go test -bench=. -cpuprofile=cpu.prof
```

### Common Issues
1. **Permission denied**: Используйте `t.TempDir()`
2. **File not found**: Проверьте относительные пути
3. **Race conditions**: Добавьте `-race` флаг
4. **Memory leaks**: Используйте benchmark тесты

---

## 📖 Knowledge Base

### Внутренние ресурсы
- **Test Design**: `_bmad-output/test-artifacts/test-design-qa.md`
- **Architecture**: `_bmad-output/test-artifacts/test-design-architecture.md`
- **Risk Matrix**: Test Design документ

### External ссылки
- [Go Testing Documentation](https://golang.org/pkg/testing/)
- [Testify Library](https://github.com/stretchr/testify)
- [Go Benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)

---

**Generated by:** BMAD TEA Framework  
**Version:** Go Test v1.0  
**Last Updated:** 2026-03-08
