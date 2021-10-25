package wrapper

import (
	"testing"
)

func FmtTest(i int) {
	panic(i)
}

func TestGo(t *testing.T) {
	//ctx := context.TODO()
	//ret := make(chan struct{})
	//Go(ctx, ret, func(ctx context.Context, ret chan<- interface{}) {
	//	s := make(chan struct{})
	//	cycle:
	//	for true {
	//		select {
	//		case <-ctx.Done():
	//			break cycle
	//		case :
	//			fmt.Println("aaa")
	//		}
	//	}
	//	t.Log("finished")
	//})
	//time.Sleep(20*time.Second)
}
