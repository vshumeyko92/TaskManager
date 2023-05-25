Менеджер задач

Менеджер задач - это простое веб-приложение, разработанное на языке Go с использованием базы данных Tarantool. Он позволяет создавать, просматривать, обновлять и удалять задачи.

Установка

Убедитесь, что у вас установлен Go и Tarantool.

Склонируйте репозиторий с помощью следующей команды:
git clone https://github.com/your-username/task-manager.git

Перейдите в директорию проекта:
cd task-manager

Установите зависимости, выполнив команду:
go get

Запустите приложение, выполнив команду:
go run main.go

Приложение будет доступно по адресу http://localhost:8080.

Использование

Создание задачи: Чтобы создать новую задачу, заполните форму на главной странице и нажмите кнопку "Создать задачу". Задача будет добавлена в список задач.
Обновление задачи: Чтобы обновить задачу, нажмите на кнопку "Редактировать" рядом с задачей в списке. Измените необходимые поля и нажмите "Сохранить".
Удаление задачи: Чтобы удалить задачу, нажмите на кнопку "Удалить" рядом с задачей в списке.

Структура проекта

main.go: Главный файл приложения, который содержит основную логику и точку входа.
task.go: Модель и функции для работы с задачами.
menu.go: Модель и функции для работы с меню приложения.
public/: Статические файлы (HTML, CSS, JavaScript) для пользовательского интерфейса.
tests/: Тестовые файлы для юнит-тестирования компонентов приложения.

Зависимости

github.com/tarantool/go-tarantool: Библиотека для работы с базой данных Tarantool.
github.com/stretchr/testify: Библиотека для написания тестов.


Чтобы запустить тесты, выполните команду go test в корневой директории проекта "task-manager". Go автоматически найдет и выполнит все тестовые функции в файлах с префиксом "Test" и выводит результаты.
