package wrapper

import (
	"context"
	"fmt"
	"time"
)

// Go 包装go关键字，限制必须指定传递ctx
// 最大执行时长10*60S
// 执行fun必须支持ctx取消
func Go(ctx context.Context, ret chan<- interface{}, fun func(ctx context.Context, ret chan<- interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Go panic：", err)
			}
		}()
		ctxc, cancel := context.WithCancel(ctx)
		time.AfterFunc(10*time.Minute, cancel)
		fun(ctxc, ret)
	}()
}
