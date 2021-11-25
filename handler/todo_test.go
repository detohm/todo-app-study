package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/detohm/todo-app-study/todo/mock"
	"github.com/golang/mock/gomock"

	"github.com/detohm/todo-app-study/todo"
	"github.com/labstack/echo/v4"
)

func TestToDoHandler_GetToDoList(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	mockService := mock.NewMockService(mockCtrl)

	e := echo.New()

	type fields struct {
		service todo.Service
	}

	tests := []struct {
		name           string
		fields         fields
		wantStatusCode int
	}{
		// TODO: Add test cases.
		{
			"normal case",
			fields{mockService},
			http.StatusOK,
		},
		{
			"service error case",
			fields{mockService},
			http.StatusInternalServerError,
		},
	}

	// expected mock service
	// normal case
	mockService.EXPECT().GetToDoList().
		Return([]todo.ToDo{{1, "test", false, false}}, nil)
	wantJSONResponse := []string{
		`[{"id":1,"description":"test","isCompleted":false,"isDeleted":false}]` + "\n",
		`fail`}

	// service error case
	mockService.EXPECT().GetToDoList().Return(nil, errors.New("service error"))

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h := NewHandler(tt.fields.service)
			req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
			rec := httptest.NewRecorder()
			_ = h.GetToDoList(e.NewContext(req, rec))

			if tt.wantStatusCode != rec.Code {
				t.Errorf("ToDoHandler.GetToDoList() http status code get = %v, expected %v", rec.Code, tt.wantStatusCode)
			}

			comp := strings.Compare(wantJSONResponse[i], rec.Body.String())
			if comp != 0 {
				t.Errorf(
					"ToDoHandler.GetToDoList() %v :json response mismatch %v: %v %v",
					tt.name,
					comp,
					rec.Body.String(),
					wantJSONResponse[i],
				)
			}
		})
	}
}

func BenchmarkToDoHandler_GetToDoList(b *testing.B) {

	mockCtrl := gomock.NewController(b)
	mockService := mock.NewMockService(mockCtrl)
	for i := 0; i < b.N; i++ {
		mockService.EXPECT().GetToDoList().
			Return([]todo.ToDo{{1, "test", false, false}}, nil)
	}
	h := NewHandler(mockService)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	rec := httptest.NewRecorder()
	e := echo.New()

	for i := 0; i < b.N; i++ {
		h.GetToDoList(e.NewContext(req, rec))
	}
}
