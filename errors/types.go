package errors

import (
	"github.com/pkg/errors"
)

type ErrorType uint

const ErrorCustomType ErrorType = 999999

// New 创建对应类型错误
func (e ErrorType) New() error {
	return custom{
		errType:    e,
		wrapperErr: errors.New(GetMsg(e)),
	}
}

// Newm 创建一个自定义错误，支持指定msg
func (e ErrorType) Newm(msg string) error {
	return custom{
		errType:    e,
		wrapperErr: errors.New(msg),
	}
}

// Wrap 包装错误
func (e ErrorType) Wrap(err error) error {
	return e.Wrapf(err, GetMsg(e))
}

// Wrapf 格式化包装错误
func (e ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	wrapErr := errors.Wrapf(err, msg, args...)
	return custom{
		errType:    e,
		wrapperErr: wrapErr,
	}
}
