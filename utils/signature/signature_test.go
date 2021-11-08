package signature

import (
	"testing"
)

type testData1 struct {
	E string `json:"e"`
	F int    `json:"f"`
}

type testData struct {
	A string    `json:"a"`
	B string    `json:"b"`
	C int       `json:"c"`
	D testData1 `json:"d"`
}

func TestSignature(t *testing.T) {
	data := testData{
		A: "abc",
		B: "你好",
		C: 19,
		D: testData1{
			E: "efg",
			F: 20,
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
