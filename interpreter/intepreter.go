package interpreter

import "io"

// Interpreter describes short interpreter interface
type Interpreter interface {
	Run() error
}

// Implementation implements Interpreter interface
type Implementation struct {
	globalVars, localVars *VarScope
	internalFuncs, globalFuncs *FuncScope
}

// Run starts execution of script
func (i *Implementation) Run() error {
	return i.globalFuncs.Execute("init", i.globalVars, i.localVars, i.internalFuncs, i.globalFuncs, 0)
}

// NewInterpreter creates new instance of Implementation object
func NewInterpreter(globalVars *VarScope, globalFuncs *FuncScope, output io.Writer, maxDepth int) (Interpreter) {
	localVars := NewVarScope()
	internalFuncs := NewInternalFuncScope(output, maxDepth)
	return &Implementation{
		globalVars: globalVars,
		localVars: localVars,
		internalFuncs: internalFuncs,
		globalFuncs: globalFuncs,
	}
}