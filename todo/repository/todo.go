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
			todo
		WHERE 
		    is_deleted = false`)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer rows.Close()

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

func (r *toDoRepository) CreateToDo(t todo.ToDo) (int64, error) {

	res, err := r.db.Exec(`
	INSERT INTO 
		todo
	(description, is_completed, is_deleted)
	VALUES (?,?,?)
	`,
		t.Description,
		false,
		false)

	if err != nil {
		fmt.Print(err)
		return -1, err
	}

	return res.LastInsertId()
}

func (r *toDoRepository) CompleteToDo(id int64) error {
	_, err := r.db.Exec(`
	UPDATE 
	   todo
	SET 
	  is_completed = true
	WHERE
	  id = ?
	`, id)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil

}

func (r *toDoRepository) DeleteToDo(id int64) error {

	_, err := r.db.Exec(`
	UPDATE 
	   todo
	SET 
	  is_deleted = true
	WHERE
	  id = ?
	`, id)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
