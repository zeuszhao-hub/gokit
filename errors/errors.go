package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type custome struct {
	errType    ErrorType
	wrapperErr error
}

// Error implementation error interface
func (c custome) Error() string {
	return c.wrapperErr.Error()
}

// Cause 实现errors.cause中的匿名接口
func (c custome) Cause() error {
	return c.wrapperErr
}

// Is 实现errors.is中的匿名接口
func (c custome) Is(target error) bool {
	if customeErr, ok := target.(custome); ok {
		return customeErr.errType == c.errType
	}
	return false
}

// Stack return call stack
func Stack(err error) string {
	if customeErr, ok := err.(custome); ok {
		return fmt.Sprintf("%+v", customeErr.wrapperErr)
	}
	return fmt.Sprintf("%+v", errors.WithStack(err))
}

// Cause return originally cause
func Cause(err error) error {
	return errors.Cause(err)
}

// Is compare source and target
func Is(source error, target error) bool {
	return errors.Is(source, target)
}
