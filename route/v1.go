package route

import (
	"github.com/detohm/todo-app-study/handler"
	"github.com/labstack/echo/v4"
)

type RouteV1 struct {
	echo        *echo.Echo
	toDoHandler *handler.ToDoHandler
}

func NewRouteV1(e *echo.Echo, h *handler.ToDoHandler) *RouteV1 {
	return &RouteV1{e, h}
}

func (r *RouteV1) Bind() {
	v1 := r.echo.Group("/api/v1")
	v1.GET("/todos", r.toDoHandler.GetToDoList)
	v1.POST("/todo", r.toDoHandler.CreateToDo)
	v1.PATCH("/todo/:id/complete", r.toDoHandler.CompleteToDo)
	v1.DELETE("/todo/:id", r.toDoHandler.DeleteTodo)
}
