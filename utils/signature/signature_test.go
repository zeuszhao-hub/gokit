package signature

import (
	"testing"
)

type testData1 struct {
	e string
	f int
}

type testData struct {
	a string
	b string
	c int
	d testData1
}

func TestSignature(t *testing.T) {
	data := testData{
		a: "abc",
		b: "bcd",
		c: 19,
		d: testData1{
			e: "efg",
			f: 20,
		},
	}

	sign, err := New("aaa", "bbb").Salt("vvv").Sign(data)
	if err != nil {
		t.Error(err)
	}
	check, err := New("aaa", "bbb").Salt("vvv").Verify(data, sign)
	if err != nil {
		t.Error(err)
	}

	if check {
		t.Log("ture")
	} else {
		t.Log("false")
	}
}
