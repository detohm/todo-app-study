package todo

type Service interface {
	GetToDoList() ([]ToDo, error)
	CreateToDo(ToDo) (int64, error)
	CompleteToDo(int64) error
	DeleteToDo(int64) error
}
