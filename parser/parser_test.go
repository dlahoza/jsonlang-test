package parser

import (
	"io"
	"testing"

	"bytes"
	"fmt"

	"github.com/DLag/jsonlang-test/interpreter"
	"github.com/stretchr/testify/assert"
)

var errMock = fmt.Errorf("some error")

type errorReaderMock struct{}

func (errorReaderMock) Read(p []byte) (n int, err error) {
	return 0, errMock
}

func TestParse_Errors(t *testing.T) {
	asserts := assert.New(t)
	cases := []struct {
		name  string
		input io.Reader
	}{
		{
			"Reader is nil",
			nil,
		},
		{
			"ReadAll error",
			&errorReaderMock{},
		},
		{
			"JSON Umarshal error",
			bytes.NewBufferString("asdasd"),
		},
		{
			"JSON unexpected type at root",
			bytes.NewBufferString(`{"a":{"b":"c"}}`),
		}, {
			"JSON unexpected type at function",
			bytes.NewBufferString(`{"a":[{"b":{}}]}`),
		},
	}
	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			globalVars := interpreter.NewVarScope()
			globalFuncs := interpreter.NewFuncScope(0)
			err := Parse(cases[i].input, globalVars, globalFuncs)
			asserts.Error(err)
		})
	}
}
