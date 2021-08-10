package errors

import (
	"github.com/pkg/errors"
)

type ErrorType uint

// New 创建对应类型错误
func (e ErrorType) New() error {
	return custome{
		errType:    e,
		wrapperErr: errors.New(getMsg(e)),
	}
}

// Wrap 包装错误
func (e ErrorType) Wrap(err error) error {
	return e.Wrapf(err, getMsg(e))
}

// Wrapf 格式化包装错误
func (e ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	wrapErr := errors.Wrapf(err, msg, args...)
	return custome{
		errType:    e,
		wrapperErr: wrapErr,
	}
}
