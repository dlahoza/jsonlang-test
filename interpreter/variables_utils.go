package interpreter

import (
	"github.com/pkg/errors"
	//"fmt"
	"strconv"
)

func getIDValue(localVars *VarScope) (id, value string, err error) {
	if id, err = localVars.Get("id"); err != nil {
		err = errors.Wrap(err, "Can't get id of variable")
		return
	}
	if value, err = localVars.Get("value"); err != nil {
		err = errors.Wrapf(err, "Can't get value of variable %q", id)
	}
	return
}

func getIDOperands(localVars *VarScope) (id, operand1, operand2 string, err error) {
	if id, _ = localVars.Get("id"); err != nil && err != ErrorVariableDoesNotExist {
		err = errors.Wrap(err, "Can't get value from local scope")
		return
	}
	if operand1, _ = localVars.Get("operand1"); err != nil && err != ErrorVariableDoesNotExist {
		err = errors.Wrap(err, "Can't get value from local scope")
		return
	}
	if operand2, _ = localVars.Get("operand2"); err != nil && err != ErrorVariableDoesNotExist {
		err = errors.Wrap(err, "Can't get value from local scope")
		return
	}
	return
}

func getIDOperandsNum(localVars *VarScope) (id string, o1, o2 float64, err error) {
	var operand1, operand2 string
	if id, operand1, operand2, err = getIDOperands(localVars); err != nil {
		return
	}
	var err1, err2 error
	o1, err1 = strconv.ParseFloat(operand1, 64)
	o2, err2 = strconv.ParseFloat(operand2, 64)
	if err1 != nil || err2 != nil {
		err = ErrorCannotConvertToFloat
	}
	return
}

func resolveVariable(value string, globalVars, localVars *VarScope) (res string, err error) {
	res = value
	if len(value) >= 2 {
		switch value[0] {
		case '#':
			id := value[1:]
			res, _ = globalVars.Get(id)
		case '$':
			id := value[1:]
			res, _ = localVars.Get(id)
		}
	}
	//fmt.Printf("Variable resolving. Input %q. Output %q.\n", value, res)
	return
}
