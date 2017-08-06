package main

import (
	"bytes"
	"fmt"
	"github.com/DLag/jsonlang-test/interpreter"
	"github.com/DLag/jsonlang-test/parser"
	"os"
)

var script = `{
  "var1":1,
  "var2":2,

  "init": [
    {"cmd" : "#setup" }
  ],

  "setup": [
    {"cmd":"update", "id": "var1", "value":3.5},
    {"cmd":"print", "value": "#var1"},
    {"cmd":"#sum", "id": "var1", "value1":"#var1", "value2":"#var2"},
    {"cmd":"print", "value": "#var1"},
    {"cmd":"create", "id": "var3", "value":5},
    {"cmd":"delete", "id": "var1"},
    {"cmd":"sub", "id": "var2", "operand1":"#var2", "operand2":"1"},
    {"cmd":"#printAll"}
  ],

  "sum": [
      {"cmd":"add", "id": "$id", "operand1":"$value1", "operand2":"$value2"}
  ],

  "printAll":
  [
    {"cmd":"print", "value": "#var1"},
    {"cmd":"print", "value": "#var2"},
    {"cmd":"print", "value": "#var3"}
  ]
}`

func main() {
	var err error
	buf := bytes.NewBufferString(script)
	globalVars := interpreter.NewVarScope()
	localVars := interpreter.NewVarScope()
	internalFuncs := interpreter.NewInternalFuncScope(10)
	globalFuncs := interpreter.NewFuncScope(10)
	if err = parser.Parse(buf, globalVars, globalFuncs); err != nil {
		fmt.Println("Error parsing script: ", err)
		os.Exit(1)
	}
	if err = globalFuncs.Execute("init", globalVars, localVars, internalFuncs, globalFuncs, 0); err != nil {
		fmt.Println("Error parsing script: ", err)
		os.Exit(1)
	}
}
