package main

import (
	"log"
	"github.com/tarantool/go-tarantool"
	"./task"
	"./menu"
	"net/http"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("public")))
}

func main() {
	// Устанавливаем соединение с Tarantool
	opts := tarantool.Opts{
		User: "guest",
	}
	conn, err := tarantool.Connect("localhost:3301", opts)
	if err != nil {
		log.Fatal("Failed to connect to Tarantool:", err)
	}
	defer conn.Close()

	// Инициализируем модуль для работы с задачами
	taskMod := &taskModule{
		conn: conn,
	}

	// Отображаем главное меню
	menu.showMenu(taskMod)
}
