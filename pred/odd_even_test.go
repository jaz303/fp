package pred

import "testing"

type customInt int

func TestOdd(t *testing.T) {
	if Odd(customInt(10)) {
		t.Fail()
	}
	if !Odd(customInt(11)) {
		t.Fail()
	}
}

func TestEven(t *testing.T) {
	if Even(customInt(11)) {
		t.Fail()
	}
	if !Even(customInt(10)) {
		t.Fail()
	}
}
