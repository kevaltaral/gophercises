package db

import (
	"testing"
)

var id int
var err error

func TestInit(t *testing.T) {
	err := Init("/home/keval/tasks.db")
	if err != nil {
		t.Errorf("ERROR(in TestInit()):%s.............", err)
	}
}

func TestCreateTask(t *testing.T) {
	_, err := CreateTask("new test task")
	if err != nil {
		t.Errorf("ERROR(in TestCreateTask()):%s.............", err)
	}
}

//func DeleteTask(key int) error

func TestDeleteTask(t *testing.T) {
	err := DeleteTask(1)
	if err != nil {
		t.Errorf("ERROR(in TestDeleteTask()):%s.............", err)

	}

}

func TestAllTasks(t *testing.T) {
	test, _ := AllTasks()
	if test == nil {

	}

}

func TestNInit(t *testing.T) {
	err := Init("/home/kl/tasks.db")
	if err == nil {
		t.Errorf("%s.............", err)
	}
}

func TestA(t *testing.T) {
	Init("/home/keval/task.db")
	DBClose()
}
