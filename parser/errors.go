package parser

import "github.com/pkg/errors"

var (
	// ErrorParsingUnexpectedType "Cannot parse, unexpected type"
	ErrorParsingUnexpectedType = errors.New("Cannot parse, unexpected type")
	// ErrorParsingNilAsReader "Reader should not be nil"
	ErrorParsingNilAsReader    = errors.New("Reader should not be nil")
)
