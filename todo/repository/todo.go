package repository

import (
	"database/sql"
	"fmt"

	"github.com/detohm/todo-app-study/todo"
)

type toDoRepository struct {
	db *sql.DB
}

func NewToDoRepository(db *sql.DB) todo.Repository {
	return &toDoRepository{db}
}

func (r *toDoRepository) GetToDoList() ([]todo.ToDo, error) {

	rows, err := r.db.Query(`
		SELECT 
			id,
			description,
			is_completed,
			is_deleted 
		FROM 
			todo`)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	results := []todo.ToDo{}

	for rows.Next() {
		var item todo.ToDo
		rows.Scan(
			&item.ID,
			&item.Description,
			&item.IsCompleted,
			&item.IsDeleted)

		results = append(results, item)
	}

	return results, nil
}
