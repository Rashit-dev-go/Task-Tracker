# Task Tracker

Простой консольный трекер задач для личного использования, написанный на Go.

## Возможности

- ✅ Добавление задач
- 📋 Просмотр списка задач
- ✏️ Отметка задач как выполненных
- 🗑️ Удаление задач
- 💾 Локальное хранение в JSON файле
- 🏠 Данные хранятся в домашней директории `~/.task-tracker/`

## Установка и использование

### Сборка из исходников

```bash
# Клонирование репозитория
git clone <repository-url>
cd task-tracker

# Сборка
make build

# Установка в систему (опционально)
make install
```

### Быстрый старт

```bash
# Инициализация хранилища
make init

# Добавление задачи
./build/task-tracker add

# Просмотр списка задач
./build/task-tracker list

# Отметить задачу как выполненную
./build/task-tracker complete <task-id>

# Удаление задачи
./build/task-tracker delete <task-id>
```

## Команды

### `init`
Инициализирует хранилище задач в `~/.task-tracker/tasks.json`

```bash
task-tracker init
```

### `add` (или `a`)
Добавляет новую задачу. Интерактивно запрашивает название и описание.

```bash
task-tracker add
```

### `list` (или `l`)
Показывает список всех задач с их статусами.

```bash
task-tracker list
```

### `complete` (или `c`)
Отмечает задачу как выполненную по её ID.

```bash
task-tracker complete <task-id>
```

### `delete` (или `d`)
Удаляет задачу по её ID.

```bash
task-tracker delete <task-id>
```

## Статусы задач

- ⭕ **todo** - задача не начата
- 🔄 **in-progress** - задача в процессе (зарезервировано для будущих версий)
- ✅ **done** - задача выполнена

## Make команды

- `make build` - собрать приложение
- `make run` - запустить с помощью флага --help
- `make install` - установить в `/usr/local/bin`
- `make clean` - очистить сборочные файлы
- `make test` - запустить тесты
- `make fmt` - форматировать код
- `make vet` - проверить код
- `make init` - инициализировать хранилище
- `make demo` - демонстрация работы
- `make check` - полная проверка (fmt + vet + test)

## Структура проекта

```
task-tracker/
├── main.go       # Основной файл с CLI интерфейсом
├── models.go     # Модели данных (Task, Status)
├── storage.go    # Работа с хранилищем JSON
├── commands.go   # Реализация команд CLI
├── go.mod        # Go модуль
├── Makefile      # Сборочные команды
└── README.md     # Документация
```

## Требования

- Go 1.21 или выше
- Linux (проверено на Ubuntu/Debian)

## Лицензия

Для личного использования.
