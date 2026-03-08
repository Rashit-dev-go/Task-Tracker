package testutils

import (
	"os"
	"path/filepath"
	"testing"
)

// SetupTempDir создает временную директорию для тестов
func SetupTempDir(t *testing.T) string {
	tempDir := t.TempDir()
	t.Logf("Created temp directory: %s", tempDir)
	return tempDir
}

// SetupTaskTrackerDir создает директорию task-tracker в указанной папке
func SetupTaskTrackerDir(t *testing.T, baseDir string) string {
	taskDir := filepath.Join(baseDir, ".task-tracker")
	err := os.MkdirAll(taskDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create task tracker directory: %v", err)
	}
	return taskDir
}

// CleanupTempDir очищает временную директорию
func CleanupTempDir(t *testing.T, tempDir string) {
	err := os.RemoveAll(tempDir)
	if err != nil {
		t.Logf("Warning: failed to cleanup temp directory %s: %v", tempDir, err)
	}
}

// CreateTestDataFile создает тестовый JSON файл с задачами
func CreateTestDataFile(t *testing.T, filePath string, content []byte) {
	err := os.WriteFile(filePath, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test data file %s: %v", filePath, err)
	}
}

// ReadTestDataFile читает тестовый JSON файл
func ReadTestDataFile(t *testing.T, filePath string) []byte {
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read test data file %s: %v", filePath, err)
	}
	return content
}

// AssertFileExists проверяет существование файла
func AssertFileExists(t *testing.T, filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Expected file %s to exist, but it doesn't", filePath)
	}
}

// AssertFileNotExists проверяет отсутствие файла
func AssertFileNotExists(t *testing.T, filePath string) {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Errorf("Expected file %s to not exist, but it does", filePath)
	}
}

// AssertDirExists проверяет существование директории
func AssertDirExists(t *testing.T, dirPath string) {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		t.Errorf("Expected directory %s to exist, but it doesn't", dirPath)
	} else if !info.IsDir() {
		t.Errorf("Expected %s to be a directory, but it's a file", dirPath)
	}
}
