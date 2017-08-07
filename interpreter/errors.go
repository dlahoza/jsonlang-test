package interpreter

import "errors"

var (
	// ErrorVariableExists "Variable already exists"
	ErrorVariableExists = errors.New("Variable already exists")
	// ErrorVariableDoesNotExist "Variable doesn't exist"
	ErrorVariableDoesNotExist = errors.New("Variable doesn't exist")
	// ErrorFunctionDoesNotExist "Function doesn't exist"
	ErrorFunctionDoesNotExist = errors.New("Function doesn't exist")
	// ErrorMaxDepth "Maximum depth is reached"
	ErrorMaxDepth = errors.New("Maximum depth is reached")
	// ErrorUserFunctionWrongCmd "Instructions in user functions should have valid "cmd" field"
	ErrorUserFunctionWrongCmd = errors.New("Instructions in user functions should have valid \"cmd\" field")
	// ErrorCannotConvertToFloat "Operands should be numeric"
	ErrorCannotConvertToFloat = errors.New("Operands should be numeric")
)
