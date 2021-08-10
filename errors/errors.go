package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type custom struct {
	errType    ErrorType
	wrapperErr error
}

// Error implementation error interface
func (c custom) Error() string {
	return c.wrapperErr.Error()
}

// Cause 实现errors.cause中的匿名接口
func (c custom) Cause() error {
	return c.wrapperErr
}

// Is 实现errors.is中的匿名接口
func (c custom) Is(target error) bool {
	if customeErr, ok := target.(custom); ok {
		return customeErr.errType == c.errType
	}
	return false
}

// Stack return call stack
func Stack(err error) string {
	if customeErr, ok := err.(custom); ok {
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

// IsCustom 判断err是否为自定义错误类型 err = nil return false
func IsCustom(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(custom)
	return ok
}
