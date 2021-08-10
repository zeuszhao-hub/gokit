package wrapper

import (
	"fmt"
)

func Go(fun func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Go panic：", err)
			}
		}()
		fun()
	}()
}
