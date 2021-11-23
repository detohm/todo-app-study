package handler

import (
	"net/http"

	"github.com/detohm/todo-app-study/todo"
	"github.com/labstack/echo/v4"
)

type ToDoHandler struct {
	service todo.Service
}

func NewHandler(service todo.Service) *ToDoHandler {
	return &ToDoHandler{service}
}

func (h *ToDoHandler) GetToDoList(c echo.Context) error {
	todos, err := h.service.GetToDoList()
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail")
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *ToDoHandler) CreateToDo(c echo.Context) error {
	t := todo.ToDo{}
	c.Bind(&t)

	insertId, err := h.service.CreateToDo(t)
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail")
	}
	return c.JSON(http.StatusOK, insertId)
}
