package helpers

import (
	"os"
	"testing"
)

var context = `
a-z
`

func TestFileDirExists(t *testing.T) {
	f, err := os.Create("testfile.conf")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(context)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove("testfile.conf")

	r := FileDirExists("testfile.conf")

	if !r {
		t.Fatal("testfile.conf should exist")
	}
}
