package primitive

import (
	"os"
	"testing"
)

func TestTransform(t *testing.T) {
	file, err := os.Open("../index.jpeg")
	if err != nil {
		t.Error("error in opening file")
	}
	_, err = Transform(file, "png", 20, WithMode(ModeCircle))
	if err != nil {
		t.Errorf("in testcase")
	}
}

//func tempfile(prefix, ext string) (*os.File, error)
//func Testtempfile()
