package repository

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/detohm/todo-app-study/todo"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("couldn't create sql mock")
	}
	return db, mock
}

func Test_toDoRepository_GetToDoList(t *testing.T) {

	db, mock := NewMock()
	repo := toDoRepository{db}
	defer func() {
		db.Close()
	}()

	// #1

	columns := []string{"id", "description", "is_completed", "is_deleted"}

	rows := sqlmock.NewRows(columns).AddRow(1, "mockdesc", false, false)

	mock.ExpectQuery(
		`SELECT id, description, is_completed, is_deleted FROM todo 
		WHERE is_deleted = false`).
		WillReturnRows(rows)

	actual, err := repo.GetToDoList()
	expected := []todo.ToDo{{1, "mockdesc", false, false}}

	if err != nil {
		t.Errorf("toDoRepository.GetToDoList() Err: %v", err)
	}

	t.Run("Verify happy path", func(t *testing.T) {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("toDoRepository.GetToDoList() = %v, want %v", actual, expected)
		}
	})

	// #2

	mock.ExpectQuery(
		`SELECT id, description, is_completed, is_deleted FROM todo 
		WHERE is_deleted = false`).
		WillReturnError(fmt.Errorf("mysql server error!"))

	t.Run("Verify error path", func(t *testing.T) {

		actual, err = repo.GetToDoList()
		if err == nil {
			t.Errorf("toDoRepository.GetToDoList() Error case : expect error but got nil")
		}
		if len(actual) > 0 {
			t.Errorf("toDoRepository.GetToDoList() Error case : expect empty slice")
		}
	})
}
