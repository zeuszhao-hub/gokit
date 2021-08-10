package errors

var msgMap map[ErrorType]string = map[ErrorType]string{
	ErrorCustomType: "系统错误",
}

// AddMsgMap 注册错误信息
func AddMsgMap(err ErrorType, msg string) {
	msgMap[err] = msg
}

// GetMsg 获取错误信息
func GetMsg(err ErrorType) string {
	val, ok := msgMap[err]
	if ok {
		return val
	}
	return msgMap[ErrorCustomType]
}

// GetCode 获取错误码
func GetCode(err error) ErrorType {
	if customeErr, ok := err.(custome); ok {
		return customeErr.errType
	}
	return ErrorCustomType
}
