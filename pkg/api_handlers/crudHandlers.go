package api_handlers

import (
	"net/http"
	"strconv"

	"example.com/todo-app/database"
	"example.com/todo-app/models"
	"github.com/gin-gonic/gin"
)

// @Tags CRUD TODOS
// @Router /todos/all [get]
// @Summary get all available todos
// @Schemes
// @Description Return an array of todos objects in json
// @Produce json
// @Success 200 {array} models.Todo
func GetTodos(c *gin.Context) {
	todos := database.GetTodos()
	c.JSON(http.StatusAccepted, todos)
}

// @Tags CRUD TODOS
// @Router /todos/{id} [get]
// @Param id path int true "Todo ID"
// @Summary get a todo by id
// @Schemes
// @Description Return a todo object in json
// @Produce json
// @Success 200 {array} models.Todo
func GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	todo, err := database.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusAccepted, todo)
}

// @Tags CRUD TODOS
// @Router /todos/new [post]
// @Param todo body models.TodoPostBody true "Todo object"
// @Summary create a new todo
// @Schemes
// @Description Create a new todo object
// @Accept json
// @Produce json
// @Success 201 {object} models.todoResponse
func CreateTodo(c *gin.Context) {
	todo := models.Todo{}
	err := c.ShouldBind(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTodo(&todo)
	response := gin.H{
		"message": "Todo created successfully",
		"todo":    todo,
	}
	c.JSON(http.StatusCreated, response)
}

// @Tags CRUD TODOS
// @Router /todos/{id} [put]
// @Param todo body models.TodoPostBody true "Todo object"
// @Summary edit a todo
// @Schemes
// @Description Edit a todo object
// @Accept json
// @Produce json
// @Success 204 {object} models.todoResponse
func EditTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	_, err = database.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	todo := models.Todo{}
	err = c.ShouldBind(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.ID = id
	database.EditTodoById(&todo)
	response := gin.H{
		"message": "Todo edited successfully",
		"todo":    todo,
	}
	c.JSON(http.StatusAccepted, response)

}
