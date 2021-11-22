package service

import (
	"github.com/detohm/todo-app-study/todo"
)

type toDoService struct {
	todoRepository todo.Repository
}

func NewToDoService(repo todo.Repository) todo.Service {
	return &toDoService{repo}
}

func (s *toDoService) GetToDoList() ([]todo.ToDo, error) {
	return s.todoRepository.GetToDoList()
}
