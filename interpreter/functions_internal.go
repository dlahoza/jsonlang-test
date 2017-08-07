package interpreter

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

// funcCreate creates variable in scope
type funcCreate struct{}

// Execute executes funcCreate logic
// Needs id and value variables in local scope
func (funcCreate) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id, value string
	if id, value, err = getIDValue(localVars); err != nil {
		err = errors.Wrap(err, "Can't get variable info")
	} else if err = globalVars.Create(id, value); err != nil {
		err = errors.Wrapf(err, "Can't create variable %q", id)
	}
	return
}

// funcUpdate updates variable in scope
type funcUpdate struct{}

// Execute executes funcUpdate logic
// Needs id and value variables in local scope
func (funcUpdate) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id, value string
	if id, value, err = getIDValue(localVars); err != nil {
		err = errors.Wrap(err, "Can't get variable info")
	} else if err = globalVars.Update(id, value); err != nil {
		err = errors.Wrapf(err, "Can't update variable %q", id)
	}
	return
}

// funcDelete deletes variable from scope
type funcDelete struct{}

// Execute executes funcDelete logic
// Needs id variable in local scope
func (funcDelete) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id string
	if id, err = localVars.Get("id"); err != nil {
		err = errors.Wrap(err, "Can't get value from local scope")
	} else {
		err = globalVars.Delete(id)
	}
	return
}

// funcPrint prints variable from scope
type funcPrint struct{}

// Execute executes funcPrint logic
// Needs id variable in local scope
func (funcPrint) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var value string
	if value, _ = localVars.Get("value"); err != nil && err != ErrorVariableDoesNotExist {
		err = errors.Wrap(err, "Can't get value from local scope")
	} else {
		fmt.Println(value)
	}
	return
}

// funcAdd adds two variables and saves result to some variable of scope
// Can concatenate strings
type funcAdd struct{}

// Execute executes funcAdd logic
// Needs id, operand1 and operand2 variables in local scope
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

// funcSub substracks two variables and saves result to some variable of scope
type funcSub struct{}

// Execute executes funcSub logic
// Needs id, operand1 and operand2 variables in local scope
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

// funcMul multiplies two variables and saves result to some variable of scope
type funcMul struct{}

// Execute executes funcMul logic
// Needs id, operand1 and operand2 variables in local scope
func (funcMul) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id, result string
	var o1, o2 float64
	if id, o1, o2, err = getIDOperandsNum(localVars); err != nil {
		err = errors.Wrap(err, "Can't get variable info")
	}
	result = strconv.FormatFloat(o1*o2, 'f', -1, 64)
	globalVars.Set(id, result)
	return
}

// funcDiv divides two variables and saves result to some variable of scope
type funcDiv struct{}

// Execute executes funcAdd logic
// Needs id, operand1 and operand2 variables in local scope
func (funcDiv) Execute(globalVars, localVars *VarScope, _, _ *FuncScope, _ int) (err error) {
	var id, result string
	var o1, o2 float64
	if id, o1, o2, err = getIDOperandsNum(localVars); err != nil {
		err = errors.Wrap(err, "Can't get variable info")
	}
	result = strconv.FormatFloat(o1/o2, 'f', -1, 64)
	globalVars.Set(id, result)
	return
}

// NewInternalFuncScope creates FuncScope and fills it with internal functions
func NewInternalFuncScope(maxDepth int) *FuncScope {
	scope := NewFuncScope(maxDepth)
	scope.Set("create", &funcCreate{})
	scope.Set("update", &funcUpdate{})
	scope.Set("delete", &funcDelete{})
	scope.Set("print", &funcPrint{})
	scope.Set("add", &funcAdd{})
	scope.Set("sub", &funcSub{})
	scope.Set("mul", &funcMul{})
	scope.Set("div", &funcDiv{})
	return scope
}
