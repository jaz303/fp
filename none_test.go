package fp

import "testing"

func TestNoneTrue(t *testing.T) {
	res := None(func(i int) bool { return i > 5 }, []int{0, -1, 2, 3})
	if !res {
		t.Fail()
	}
}

func TestNoneFalse(t *testing.T) {
	res := None(func(i int) bool { return i > 5 }, []int{-1, -2, 10, 4})
	if res {
		t.Fail()
	}
}

func TestNonePtrTrue(t *testing.T) {
	res := NonePtr(func(i *int) bool { return *i > 5 }, []int{0, -1, 2, 3})
	if !res {
		t.Fail()
	}
}

func TestNonePtrFalse(t *testing.T) {
	res := NonePtr(func(i *int) bool { return *i > 5 }, []int{-1, -2, 10, 4})
	if res {
		t.Fail()
	}
}
