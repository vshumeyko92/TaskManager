package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/tarantool/go-tarantool"
)

func TestCreateTask(t *testing.T) {
	// Создаем соединение с тестовой базой данных Tarantool
	opts := tarantool.Opts{
		User: "guest",
	}
	conn, err := tarantool.Connect("localhost:3301", opts)
	if err != nil {
		t.Fatal("Failed to connect to Tarantool:", err)
	}
	defer conn.Close()

	// Инициализируем модуль для работы с задачами
	taskMod := &taskModule{
		conn: conn,
	}

	// Создаем тестовую задачу
	task := Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "Pending",
	}

	// Вызываем функцию создания задачи
	err = taskMod.createTask(&task)
	if err != nil {
		t.Fatal("Failed to create task:", err)
	}

	// Проверяем, что задача успешно создана и получила ID
	assert.NotZero(t, task.ID, "Task ID should not be zero")
}

// Добавьте другие тестовые функции для проверки остальной функциональности модуля "task"

func TestMain(m *testing.M) {
	// Здесь можно добавить код для настройки окружения тестов

	// Запуск тестов
	code := m.Run()

	// Здесь можно добавить код для очистки окружения тестов

	// Завершение
	os.Exit(code)
}
