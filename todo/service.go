package todo

type Service interface {
	GetToDoList() ([]ToDo, error)
}
