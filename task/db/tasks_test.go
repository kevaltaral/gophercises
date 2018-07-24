package db

import (
	"testing"
)

var id int
var err error

func TestInit(t *testing.T) {
	err := Init("/home/keval/tasks.db")
	if err != nil {
		t.Errorf("%s.............", err)
	}
}

func TestCreateTask(t *testing.T) {
	_, err := CreateTask("new test task")
	if err != nil {
		t.Errorf("In createtask.........")
	}
}

//func DeleteTask(key int) error

func TestDeleteTask(t *testing.T) {
	err := DeleteTask(1)
	if err != nil {
		t.Errorf("%s........", err)

	}

}
