package service

import (
	"reflect"
	"testing"

	"github.com/detohm/todo-app-study/todo/mock"
	"github.com/golang/mock/gomock"

	"github.com/detohm/todo-app-study/todo"
)

func Test_toDoService_GetToDoList(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().GetToDoList().Return([]todo.ToDo{}, nil)

	type fields struct {
		todoRepository todo.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []todo.ToDo
		wantErr bool
	}{
		{
			"blank case",
			fields{mockRepo},
			[]todo.ToDo{},
			false}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoService{
				todoRepository: tt.fields.todoRepository,
			}
			got, err := s.GetToDoList()
			if (err != nil) != tt.wantErr {
				t.Errorf("toDoService.GetToDoList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toDoService.GetToDoList() = %v, want %v", got, tt.want)
			}
		})
	}
}
