package interpreter

import (
	"fmt"
	"github.com/pkg/errors"
)

// Function - basic function interface which describes Execute method
type Function interface {
	// Execute executes logic provided by Function object
	Execute(globalVars, localVars *VarScope, internalFuncs, globalFuncs *FuncScope, depth int) error
}

// FuncScope scope of functions
type FuncScope struct {
	// maxDepth contains maximum value of execution stack depth
	maxDepth int
	funcs    map[string]Function
}

// NewFuncScope creates new scrop of functions
func NewFuncScope(maxDepth int) *FuncScope {
	return &FuncScope{
		maxDepth: maxDepth,
		funcs:    make(map[string]Function),
	}
}

// Execute executes requested function from scope and controls execution stack depth
func (s *FuncScope) Execute(name string, globalVars, localVars *VarScope, internalFuncs, globalFuncs *FuncScope, depth int) (err error) {
	errorStr := fmt.Sprintf("Can't execute function %q", name)
	depth++
	if depth > s.maxDepth {
		err = errors.Wrap(ErrorMaxDepth, errorStr)
	} else if f, ok := s.funcs[name]; !ok {
		err = errors.Wrapf(ErrorFunctionDoesNotExist, errorStr)
	} else {
		err = f.Execute(globalVars, localVars, internalFuncs, globalFuncs, depth)
	}
	return
}

// Set adds or updates function in stack by name
func (s *FuncScope) Set(name string, f Function) {
	s.funcs[name] = f
}

// UserFunction implements user-defined function which satisfy Function interface
type UserFunction struct {
	body []map[string]string
}

// Execute executes user-defined function
func (f UserFunction) Execute(globalVars, localVars *VarScope, internalFuncs, globalFuncs *FuncScope, depth int) (err error) {
	for i, instruction := range f.body {
		var cmd string
		newLocalVars := NewVarScope()
		for k, v := range instruction {
			if k == "cmd" {
				cmd = v
			} else {
				if v, err = resolveVariable(v, globalVars, localVars); err != nil {
					err = errors.Wrap(err, "Can't execute user function")
					return
				}
				newLocalVars.Set(k, v)
			}
		}
		if len(cmd) == 0 {
			err = ErrorUserFunctionWrongCmd
			return
		}
		if cmd[0] == '#' {
			if err = globalFuncs.Execute(cmd[1:], globalVars, newLocalVars, internalFuncs, globalFuncs, depth); err != nil {
				err = errors.Wrapf(err, "Error when executing global instruction #%d: %q", i, cmd)
				return
			}
		} else {
			if err = internalFuncs.Execute(cmd, globalVars, newLocalVars, internalFuncs, globalFuncs, depth); err != nil {
				err = errors.Wrapf(err, "Error when executing internal instruction #%d: %q", i, cmd)
				return
			}
		}
	}
	return
}

// NewUserFunction creates user-defined function object from instruction list
func NewUserFunction(body []map[string]string) *UserFunction {
	return &UserFunction{body}
}
