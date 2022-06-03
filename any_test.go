package fp

import "testing"

func TestAnyTrue(t *testing.T) {
	res := Any(func(i int) bool { return i > 5 }, []int{0, -1, 5, 10})
	if !res {
		t.Fail()
	}
}

func TestAnyFalse(t *testing.T) {
	res := Any(func(i int) bool { return i > 5 }, []int{-1, -2, -3, 4})
	if res {
		t.Fail()
	}
}

func TestAnyPtrTrue(t *testing.T) {
	res := AnyPtr(func(i *int) bool { return *i > 5 }, []int{0, -1, 5, 10})
	if !res {
		t.Fail()
	}
}

func TestAnyPtrFalse(t *testing.T) {
	res := AnyPtr(func(i *int) bool { return *i > 5 }, []int{-1, -2, -3, 4})
	if res {
		t.Fail()
	}
}
