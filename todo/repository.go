package todo

type Repository interface {
	GetToDoList() ([]ToDo, error)
}