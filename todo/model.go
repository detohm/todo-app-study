package todo

type ToDo struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
	IsDeleted   bool   `json:"isDeleted"`
}
