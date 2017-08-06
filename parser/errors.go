package parser

import "github.com/pkg/errors"

var (
	ErrorCannotReadProgram     = "Cannot read the program"
	ErrorParsingUnexpectedType = errors.New("Cannot parse, unexpected")
	ErrorParsingNilAsReader    = errors.New("Reader should not be nil")
)
