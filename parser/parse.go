package parser

import (
	"encoding/json"
	"github.com/DLag/jsonlang-test/interpreter"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"strconv"
)

// Parse consumes io.Reader and parses JSON FSL to internal representation
func Parse(r io.Reader, globalVars *interpreter.VarScope, globalFuncs *interpreter.FuncScope) (err error) {
	if r == nil {
		err = ErrorParsingNilAsReader
		return
	}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		err = errors.Wrap(err, "Cannot read the program")
		return
	}
	program := make(map[string]interface{})
	err = json.Unmarshal(buf, &program)
	if err != nil {
		err = errors.Wrap(err, "Cannot parse the script")
		return
	}
	for item, value := range program {
		switch value.(type) {
		case string:
			// Global variable with string value
			globalVars.Set(item, value.(string))
		case float64:
			// Global variable with number value
			globalVars.Set(item, strconv.FormatFloat(value.(float64), 'f', -1, 64))
		case []interface{}:
			// Function
			var body []map[string]string
			body, err = parseFunction(value)
			if err != nil {
				return
			}
			globalFuncs.Set(item, interpreter.NewUserFunction(body))
		default:
			err = ErrorParsingUnexpectedType
			return
		}
	}
	return
}

func parseFunction(rawFunction interface{}) (body []map[string]string, err error) {
	function := rawFunction.([]interface{})
	body = make([]map[string]string, len(function))
	for _, rawCommand := range rawFunction.([]interface{}) {
		if command, ok := rawCommand.(map[string]interface{}); ok {
			var elements map[string]string
			elements, err = parseCommand(command)
			if err != nil {
				return
			}
			body = append(body, elements)
		} else {
			err = ErrorParsingUnexpectedType
			return
		}
	}
	return
}

func parseCommand(command map[string]interface{}) (elements map[string]string, err error) {
	elements = make(map[string]string, len(command))
	for k, rawElement := range command {
		switch rawElement.(type) {
		case string:
			// Global variable with string value
			elements[k] = rawElement.(string)
		case float64:
			// Global variable with numeric value
			elements[k] = strconv.FormatFloat(rawElement.(float64), 'f', -1, 64)
		default:
			err = ErrorParsingUnexpectedType
			return
		}
	}
	return
}
