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

func (s *toDoService) CreateToDo(t todo.ToDo) (int64, error) {
	return s.todoRepository.CreateToDo(t)
}

func (s *toDoService) CompleteToDo(id int64) error {
	return s.todoRepository.CompleteToDo(id)
}

func (s *toDoService) DeleteToDo(id int64) error {
	return s.todoRepository.DeleteToDo(id)
}
