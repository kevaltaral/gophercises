package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"

	"github.com/kevaltaral/gophercises/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

/*func TestMain(t *testing.M) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)
}*/

func initial() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)

}

func TestAddCmd(t *testing.T) {
	file, _ := os.Create("./test.txt")
	defer file.Close()
	defer os.Remove(file.Name())

	old := os.Stdout
	os.Stdout = file
	initial()
	testSuit := []struct {
		args     []string
		expected string
	}{
		{args: []string{"go", "to", "gym"}, expected: `Added "go to gym"`},
		{args: []string{"go", "to", "office"}, expected: `Added "go to office"`},
		{args: []string{"clean", "dishes"}, expected: `Added "clean dishes"`},
	}
	for i, test := range testSuit {
		addCmd.Run(addCmd, test.args)
		file.Seek(0, 0)
		b, _ := ioutil.ReadFile(file.Name())
		match, err := regexp.Match(test.expected, b)

		if err != nil {
			t.Error("error in regex")
		}
		if !match {
			t.Errorf("error in add command, %d", i)
		}
	}

	os.Stdout = old

}

func TestListCLI(t *testing.T) {
	file, _ := os.Create("./test.txt")
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	//initial()
	args := []string{}
	listCmd.Run(listCmd, args)
	expected := `1. go to gym
2. go to office
3. clean dishes
`
	b, _ := ioutil.ReadFile(file.Name())
	output := string(b)
	if expected != output {
		t.Error("error in list command")
	}
	os.Stdout = old
}

func TestDoCLI(t *testing.T) {
	file, _ := os.Create("./test.txt")
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file

	//initial()
	testSuit := []struct {
		args     []string
		expected string
	}{
		{args: []string{"1", "2", "3", "a"}, expected: `Failed to parse the argument: a
Marked "1" as completed.
Marked "2" as completed.
Marked "3" as completed.
`},
		{args: []string{"1"}, expected: "Invalid task number: 1"},
	}
	for i, test := range testSuit {
		doCmd.Run(doCmd, test.args)
		file.Seek(0, 0)
		b, _ := ioutil.ReadFile(file.Name())
		match, err := regexp.Match(test.expected, b)
		if err != nil {
			t.Error("error in regex")
		}
		if !match {
			t.Errorf("error .......%d", i)
		}
	}
	os.Stdout = old
}

func TestNListCLI(t *testing.T) {
	file, _ := os.Create("./test.txt")
	defer file.Close()
	//defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	//initial()
	args := []string{}
	listCmd.Run(listCmd, args)
	file.Seek(0, 0)
	expected := "You have no tasks to complete!"
	b, _ := ioutil.ReadFile(file.Name())
	output := string(b)
	os.Stdout = old
	fmt.Printf(output)
	if reflect.DeepEqual(expected, output) {
		t.Error("error in list command")
	}
}

func TestNegative(t *testing.T) {
	db.DBClose()
	file, _ := os.Create("./test.txt")
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	//initial()
	args := []string{}
	listCmd.Run(listCmd, args)
	expected := "Something went wrong:"
	b, _ := ioutil.ReadFile(file.Name())
	output := string(b)
	if reflect.DeepEqual(expected, output) {
		t.Error("error in list command")
	}
	os.Stdout = old
}

func TestNegativeDo(t *testing.T) {
	db.DBClose()
	file, _ := os.Create("./test.txt")
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	//initial()
	args := []string{}
	doCmd.Run(listCmd, args)
	expected := "Something went wrong:"
	b, _ := ioutil.ReadFile(file.Name())
	output := string(b)
	if reflect.DeepEqual(expected, output) {
		t.Error("error in do command")
	}
	os.Stdout = old
}
