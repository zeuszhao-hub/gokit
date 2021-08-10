package wrapper

import (
	"testing"
)

func FmtTest(i int) {
	panic(i)
}

func TestGo(t *testing.T) {
	Go(func() {
		FmtTest(1)
	})
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Go(func() {
			FmtTest(i)
		})
	}
}
