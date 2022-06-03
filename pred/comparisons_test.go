package pred

import "testing"

func intPtr(v int) *int {
	return &v
}

func TestZeroNumber(t *testing.T) {
	if Zero(1) || !Zero(0) {
		t.Fail()
	}
}

func TestZeroString(t *testing.T) {
	if Zero("foo") || !Zero("") {
		t.Fail()
	}
}

func TestZeroStruct(t *testing.T) {
	type foo struct {
		a string
	}

	if Zero(foo{"hello"}) || !Zero(foo{}) {
		t.Fail()
	}
}

func TestGt(t *testing.T) {
	fn := Gt(10)
	if !fn(11) || fn(9) {
		t.Fail()
	}
}

func TestGtPtr(t *testing.T) {
	fn := GtPtr(10)
	if !fn(intPtr(11)) || fn(intPtr(9)) {
		t.Fail()
	}
}

func TestGtEq(t *testing.T) {
	fn := GtEq(10)
	if !fn(10) || fn(9) {
		t.Fail()
	}
}

func TestGtEqPtr(t *testing.T) {
	fn := GtEqPtr(10)
	if !fn(intPtr(10)) || fn(intPtr(9)) {
		t.Fail()
	}
}

func TestLt(t *testing.T) {
	fn := Lt(10)
	if fn(10) || !fn(9) {
		t.Fail()
	}
}

func TestLtPtr(t *testing.T) {
	fn := LtPtr(10)
	if fn(intPtr(10)) || !fn(intPtr(9)) {
		t.Fail()
	}
}

func TestLtEq(t *testing.T) {
	fn := LtEq(10)
	if !fn(10) || fn(11) {
		t.Fail()
	}
}

func TestLtEqPtr(t *testing.T) {
	fn := LtEqPtr(10)
	if !fn(intPtr(10)) || fn(intPtr(11)) {
		t.Fail()
	}
}

func TestEq(t *testing.T) {
	fn := Eq(10)
	if !fn(10) || fn(9) {
		t.Fail()
	}
}

func TestEqPtr(t *testing.T) {
	fn := EqPtr(10)
	if !fn(intPtr(10)) || fn(intPtr(9)) {
		t.Fail()
	}
}

func TestNotEq(t *testing.T) {
	fn := NotEq(10)
	if fn(10) || !fn(9) {
		t.Fail()
	}
}

func TestNotEqPtr(t *testing.T) {
	fn := NotEqPtr(10)
	if fn(intPtr(10)) || !fn(intPtr(9)) {
		t.Fail()
	}
}
