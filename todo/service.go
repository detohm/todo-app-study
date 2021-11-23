package todo

type Service interface {
	GetToDoList() ([]ToDo, error)
	CreateToDo(ToDo) (int64, error)
}
