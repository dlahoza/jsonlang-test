package parser

import (
	"encoding/json"
	"github.com/DLag/jsonlang-test/interpreter"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"strconv"
)

func Parse(r io.Reader, globalVars *interpreter.VarScope, globalFuncs *interpreter.FuncScope) (err error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		errors.Wrap(err, ErrorCannotReadProgram)
	}
	program := make(map[string]interface{})
	err = json.Unmarshal(buf, &program)
	if err != nil {
		errors.Wrap(err, "Cannot parse the script")
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
			for _, rawCommand := range value.([]interface{}) {
				if command, ok := rawCommand.(map[string]interface{}); ok {
					elements := make(map[string]string, len(command))
					for k, rawElement := range command {
						switch rawElement.(type) {
						case string:
							elements[k] = rawElement.(string)
						case float64:
							// Global variable with number value
							elements[k] = strconv.FormatFloat(rawElement.(float64), 'f', -1, 64)
						default:
							err = ErrorParsingUnexpectedType
							return
						}
					}
					body = append(body, elements)
				} else {
					err = ErrorParsingUnexpectedType
					return
				}
			}
			globalFuncs.Set(item, interpreter.NewUserFunction(body))
		default:
			err = ErrorParsingUnexpectedType
			return
		}
	}
	return
}
