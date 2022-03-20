package pkg

import "testing"

func Test_Sum(t *testing.T) {
	result := Sum(1, 1)

	if result != 2 {
		t.Error("should return 2 when receive a=1, b=1")
	}
}

func Test_Abs(t *testing.T) {
	result := Abs(-1)

	if result == -1 {
		t.Error("should return 1 when receive a=-1")
	}
}
