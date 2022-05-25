package util

import (
	"fmt"
)

// ThrowError create new error with default msg constrain
func ThrowError(datatype, function, msg string) error {
	return fmt.Errorf("%s.%s  -  failed, because: %s", datatype, function, msg)
}

// ForwardError update thrown error with call stack parent
func ForwardError(datatype, function string, err error) error {
	return fmt.Errorf("%s.%s  -  failed, because: \n\t%s", datatype, function, err)
}
