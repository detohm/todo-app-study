package todo

type Repository interface {
	GetToDoList() ([]ToDo, error)
	CreateToDo(ToDo) (int64, error)
}
