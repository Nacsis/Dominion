package util

import (
	"fmt"
)

type ErrorInfo struct {
	FunctionName, FileName string
}

// ThrowError create new error with default msg constrain
func (ei *ErrorInfo) ThrowError(msg string) error {
	return fmt.Errorf("%s.%s  -  failed, because: %s", ei.FileName, ei.FunctionName, msg)
}

// ForwardError update thrown error with call stack parent
func (ei *ErrorInfo) ForwardError(err error) error {
	return fmt.Errorf("%s.%s  -  failed, because: \n\t%s", ei.FileName, ei.FunctionName, err)
}
