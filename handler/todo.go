package handler

import (
	"net/http"
	"strconv"

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

func (h *ToDoHandler) CompleteToDo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail")
	}
	id64 := int64(id)
	err = h.service.CompleteToDo(id64)
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail")
	}
	return c.JSON(http.StatusOK, "ok")
}

func (h *ToDoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail")
	}
	id64 := int64(id)
	err = h.service.DeleteToDo(id64)
	if err != nil {
		return c.String(http.StatusInternalServerError, "fail")
	}
	return c.JSON(http.StatusOK, "ok")
}
