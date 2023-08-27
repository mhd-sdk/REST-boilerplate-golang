package database

import (
	"fmt"

	"example.com/todo-app/models"
)

func GetTodos() []models.Todo {
	var todos []models.Todo
	Database.Find(&todos)
	return todos
}

func GetTodoById(id int) (models.Todo, error) {
	todo := models.Todo{}
	err := Database.First(&todo, id).Error
	if err != nil {
		fmt.Println("failed to get todo")
	}
	return todo, err
}

func CreateTodo(todo *models.Todo) error {
	// cast TodoPostBody to Todo
	castedTodo := models.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
	}

	err := Database.Create(&castedTodo).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func EditTodoById(todo *models.Todo) {
	Database.Save(todo)
}
