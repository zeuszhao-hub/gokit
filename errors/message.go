package errors

var msgMap map[ErrorType]string = map[ErrorType]string{}

// AddMsgMap 注册错误信息
func AddMsgMap(err ErrorType, msg string) {
	msgMap[err] = msg
}

// GetMsg 获取错误信息
func GetMsg(err ErrorType) string {
	val, ok := msgMap[err]
	if ok != true {
		return "未注册错误信息"
	}
	return val
}
