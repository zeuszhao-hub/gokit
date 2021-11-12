package texture

import "testing"

type A struct {
	A string
}

type B struct {
	A
}

func TestTo(t *testing.T) {
	a := A{A: "abc"}
	b := B{}
	err := To(a, &b)
	if err != nil {
		t.Error(err)
	}

	t.Log(b)
}
