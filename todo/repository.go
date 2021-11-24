//go:generate mockgen -source ./repository.go -destination ./mock/repository.go -package=mock
package todo

type Repository interface {
	GetToDoList() ([]ToDo, error)
	CreateToDo(ToDo) (int64, error)
	CompleteToDo(int64) error
	DeleteToDo(int64) error
}
