package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/kevaltaral/gophercises/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func initial() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)
}
func TestAddCmd(t *testing.T) {
	file, _ := os.Create("./test.txt")
	defer file.Close()
	//defer os.Remove(file.Name())
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
