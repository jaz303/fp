package fp

import "testing"

func TestFirstIndexOfPresent(t *testing.T) {
	ix := FirstIndexOf(func(i int) bool {
		return i > 10
	}, []int{0, 5, 10, 15})
	if ix != 3 {
		t.Fail()
	}
}

func TestFirstIndexOfMissing(t *testing.T) {
	ix := FirstIndexOf(func(i int) bool {
		return i > 100
	}, []int{0, 5, 10, 15})
	if ix != -1 {
		t.Fail()
	}
}

func TestFirstIndexOfPtrPresent(t *testing.T) {
	ix := FirstIndexOfPtr(func(i *int) bool {
		return *i > 10
	}, []int{0, 5, 10, 15})
	if ix != 3 {
		t.Fail()
	}
}

func TestFirstIndexOfPtrMissing(t *testing.T) {
	ix := FirstIndexOfPtr(func(i *int) bool {
		return *i > 100
	}, []int{0, 5, 10, 15})
	if ix != -1 {
		t.Fail()
	}
}
