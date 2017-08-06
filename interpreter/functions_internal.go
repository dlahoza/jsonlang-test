package interpreter

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type funcCreate struct{}

func (funcCreate) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id, value string
	if id, value, err = getIDValue(localVars); err != nil {
		err = errors.Wrap(err, "Can't get variable info")
	} else if err = globalVars.Create(id, value); err != nil {
		err = errors.Wrapf(err, "Can't funcCreate variable %q", id)
	}
	return
}

type funcUpdate struct{}

func (funcUpdate) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id, value string
	if id, value, err = getIDValue(localVars); err != nil {
		err = errors.Wrap(err, "Can't get variable info")
	} else if err = globalVars.Update(id, value); err != nil {
		err = errors.Wrapf(err, "Can't funcUpdate variable %q", id)
	}
	return
}

type funcDelete struct{}

func (funcDelete) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id string
	if id, err = localVars.Get("id"); err != nil {
		err = errors.Wrap(err, "Can't get value from local scope")
	} else {
		globalVars.Delete(id)
	}
	return
}

type funcPrint struct{}

func (funcPrint) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var value string
	if value, _ = localVars.Get("value"); err != nil && err != ErrorVariableDoesNotExist {
		err = errors.Wrap(err, "Can't get value from local scope")
	} else {
		fmt.Println(value)
	}
	return
}

type funcAdd struct{}

func (funcAdd) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id, operand1, operand2, result string
	if id, operand1, operand2, err = getIDOperands(localVars); err != nil {
		err = errors.Wrap(err, "Can't get variable info")
	}
	o1, err1 := strconv.ParseFloat(operand1, 64)
	o2, err2 := strconv.ParseFloat(operand2, 64)
	if err1 != nil || err2 != nil {
		result = operand1 + operand2
	} else {
		result = strconv.FormatFloat(o1+o2, 'f', -1, 64)
	}
	globalVars.Set(id, result)
	return
}

type funcSub struct{}

func (funcSub) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id, result string
	var o1, o2 float64
	if id, o1, o2, err = getIDOperandsNum(localVars); err != nil {
		err = errors.Wrap(err, "Can't get variable info")
	}
	result = strconv.FormatFloat(o1-o2, 'f', -1, 64)
	globalVars.Set(id, result)
	return
}

func NewInternalFuncScope(maxDepth int) *FuncScope {
	scope := NewFuncScope(maxDepth)
	scope.Set("create", &funcCreate{})
	scope.Set("update", &funcUpdate{})
	scope.Set("delete", &funcDelete{})
	scope.Set("print", &funcPrint{})
	scope.Set("add", &funcAdd{})
	scope.Set("sub", &funcSub{})
	return scope
}
