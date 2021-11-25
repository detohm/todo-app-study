package service

import (
	"reflect"
	"testing"

	"github.com/detohm/todo-app-study/todo"
	"github.com/detohm/todo-app-study/todo/mock"
	"github.com/golang/mock/gomock"
)

func Test_toDoService_GetToDoList(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockRepository(mockCtrl)

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
			false},
		{
			"one item case",
			fields{mockRepo},
			[]todo.ToDo{{1, "case2", false, false}},
			false},
	}

	// expect mock repo by test cases
	// 1:blank case
	mockRepo.EXPECT().GetToDoList().Return([]todo.ToDo{}, nil)
	// 2:one item case
	mockRepo.EXPECT().GetToDoList().
		Return([]todo.ToDo{{1, "case2", false, false}}, nil)

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

func Test_toDoService_CreateToDo(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockRepository(mockCtrl)

	type fields struct {
		todoRepository todo.Repository
	}
	type args struct {
		t todo.ToDo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"normal case",
			fields{mockRepo},
			args{todo.ToDo{0, "normal", false, false}},
			1,
			false,
		},
	}

	// expected mock repo by test cases
	// 1:normal case
	mockRepo.EXPECT().CreateToDo(tests[0].args.t).Return(int64(1), nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &toDoService{
				todoRepository: tt.fields.todoRepository,
			}
			got, err := s.CreateToDo(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("toDoService.CreateToDo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toDoService.CreateToDo() = %v, want %v", got, tt.want)
			}
		})
	}
}
