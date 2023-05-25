package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

// Определение функций для отображения меню и обработки выбора пользователя (showMenu, handleCreateTask, handleViewTask, handleUpdateTask, handleDeleteTask, handleListTasks)

// showMenu отображает главное меню приложения
func showMenu(taskMod *taskModule) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("=== Менеджер задач ===")
		fmt.Println("1. Создать задачу")
		fmt.Println("2. Просмотреть задачу")
		fmt.Println("3. Обновить задачу")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Вывести список задач")
		fmt.Println("6. Выйти")
		fmt.Print("Выберите опцию: ")

		if !scanner.Scan() {
			fmt.Println("Ошибка чтения ввода:", scanner.Err())
			continue
		}

		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			handleCreateTask(taskMod)
		case "2":
			handleViewTask(taskMod)
		case "3":
			handleUpdateTask(taskMod)
		case "4":
			handleDeleteTask(taskMod)
		case "5":
			handleListTasks(taskMod)
		case "6":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверная опция. Повторите попытку.")
		}
	}
}

// handleCreateTask обрабатывает создание новой задачи
func handleCreateTask(taskMod *taskModule) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Создание задачи ===")
	fmt.Print("Введите название задачи: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка чтения ввода:", scanner.Err())
		return
	}
	title := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите описание задачи: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка чтения ввода:", scanner.Err())
		return
	}
	description := strings.TrimSpace(scanner.Text())

	// Создаем задачу
	task := Task{
		Title:       title,
		Description: description,
		Status:      "В процессе",
	}
	err := taskMod.createTask(task)
	if err != nil {
		fmt.Println("Ошибка при создании задачи:", err)
	} else {
		fmt.Println("Задача успешно создана.")
	}
}

// handleViewTask обрабатывает просмотр задачи по идентификатору
func handleViewTask(taskMod *taskModule) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Просмотр задачи ===")
	fmt.Print("Введите идентификатор задачи: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка чтения ввода:", scanner.Err())
		return
	}
	taskID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Ошибка ввода идентификатора задачи:", err)
		return
	}

	// Получаем задачу по идентификатору
	task, err := taskMod.getTask(taskID)
	if err != nil {
		fmt.Println("Ошибка при получении задачи:", err)
	} else {
		fmt.Println("=== Информация о задаче ===")
		fmt.Println("ID:", task.ID)
		fmt.Println("Название:", task.Title)
		fmt.Println("Описание:", task.Description)
		fmt.Println("Статус:", task.Status)
	}
}

// handleUpdateTask обрабатывает обновление задачи
func handleUpdateTask(taskMod *taskModule) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Обновление задачи ===")
	fmt.Print("Введите идентификатор задачи: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка чтения ввода:", scanner.Err())
		return
	}
	taskID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Ошибка ввода идентификатора задачи:", err)
		return
	}

	fmt.Print("Введите новое название задачи: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка чтения ввода:", scanner.Err())
		return
	}
	title := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите новое описание задачи: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка чтения ввода:", scanner.Err())
		return
	}
	description := strings.TrimSpace(scanner.Text())

	// Обновляем задачу
	updatedTask := Task{
		ID:          taskID,
		Title:       title,
		Description: description,
	}
	err = taskMod.updateTask(taskID, updatedTask)
	if err != nil {
		fmt.Println("Ошибка при обновлении задачи:", err)
	} else {
		fmt.Println("Задача успешно обновлена.")
	}
}

// handleDeleteTask обрабатывает удаление задачи
func handleDeleteTask(taskMod *taskModule) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Удаление задачи ===")
	fmt.Print("Введите идентификатор задачи: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка чтения ввода:", scanner.Err())
		return
	}
	taskID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Ошибка ввода идентификатора задачи:", err)
		return
	}

	// Удаляем задачу
	err = taskMod.deleteTask(taskID)
	if err != nil {
		fmt.Println("Ошибка при удалении задачи:", err)
	} else {
		fmt.Println("Задача успешно удалена.")
	}
}

// handleListTasks обрабатывает вывод списка всех задач
func handleListTasks(taskMod *taskModule) {
	tasks, err := taskMod.listTasks()
	if err != nil {
		fmt.Println("Ошибка при получении списка задач:", err)
		return
	}

	fmt.Println("=== Список задач ===")
	for _, task := range tasks {
		fmt.Println("ID:", task.ID)
		fmt.Println("Название:", task.Title)
		fmt.Println("Описание:", task.Description)
		fmt.Println("Статус:", task.Status)
		fmt.Println()
	}
}

// createTask создает новую задачу
func (tm *taskModule) createTask(task Task) error {
	_, err := tm.conn.Insert("tasks", []interface{}{
		task.ID,
		task.Title,
		task.Description,
		task.Status,
	})
	return err
}

// getTask возвращает задачу по идентификатору
func (tm *taskModule) getTask(taskID int) (Task, error) {
	var task Task
	err := tm.conn.SelectTyped("tasks", "primary", 0, 1, tarantool.IterEq, []interface{}{taskID}, &task)
	return task, err
}

// updateTask обновляет задачу с указанным идентификатором
func (tm *taskModule) updateTask(taskID int, updatedTask Task) error {
	_, err := tm.conn.Update("tasks", "primary", []interface{}{taskID}, []interface{}{
		[]interface{}{"=", 1, updatedTask.Title},
		[]interface{}{"=", 2, updatedTask.Description},
	})
	return err
}

// deleteTask удаляет задачу по идентификатору
func (tm *taskModule) deleteTask(taskID int) error {
	_, err := tm.conn.Delete("tasks", "primary", []interface{}{taskID})
	return err
}

// listTasks возвращает список всех задач
func (tm *taskModule) listTasks() ([]Task, error) {
	var tasks []Task
	_, err := tm.conn.SelectTyped("tasks", "primary", 0, 0, tarantool.IterEq, []interface{}{}, &tasks)
	return tasks, err
}
