package main

import (
	"github.com/tarantool/go-tarantool"
)

// Task структура задачи
type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

// taskModule модуль для работы с задачами
type taskModule struct {
	conn *tarantool.Connection
}

// Определение функций для работы с задачами (createTask, getTask, updateTask, deleteTask, listTasks)
...
