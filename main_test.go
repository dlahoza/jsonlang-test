package main

import (
	"bytes"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	oldGetFiles := getFiles
	defer func() {
		getFiles = oldGetFiles
	}()
	getFiles = func() []string {
		return []string{"examples/script1.json", "examples/script2.json"}
	}
	buf := bytes.NewBuffer(nil)
	output = buf
	defer func() {
		output = os.Stdout
	}()
	main()
	expected := "3.5\n5.5\nundefined\n1\n5\n10\n1\n5\n50\n2\nHello, world!\n"
	assert.Equal(t, expected, buf.String())
}
