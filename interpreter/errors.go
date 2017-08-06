package interpreter

import "errors"

var (
	ErrorVariableExists       = errors.New("Variable already exists")
	ErrorVariableDoesNotExist = errors.New("Variable doesn't exist")
	ErrorFunctionDoesNotExist = errors.New("Function doesn't exist")
	ErrorMaxDepth             = errors.New("Maximum depth is reached")
	ErrorUserFunctionWrongCmd = errors.New("Instructions in user functions should have valid \"cmd\" field")
	ErrorCannotConvertToFloat = errors.New("Operands should be numeric")
)
