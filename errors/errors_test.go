package errors

import (
	"github.com/pkg/errors"
	"testing"
)

const Error ErrorType = 500
const NotFound ErrorType = 404

var sentinelErr = errors.New("sentinel error")

func TestErrors(t *testing.T) {
	err1 := NotFound.Wrap(sentinelErr)
	ass := Is(err1, Error.New())
	if ass {
		t.Error("错误值断言失败")
	} else {
		t.Log("错误值断言成功")
	}
}
