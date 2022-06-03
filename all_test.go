package fp

import "testing"

func TestAllTrue(t *testing.T) {
	res := All(func(i int) bool { return i > 5 }, []int{6, 7, 8, 9})
	if !res {
		t.Fail()
	}
}

func TestAllFalse(t *testing.T) {
	res := All(func(i int) bool { return i > 5 }, []int{6, 7, 8, 9, 1})
	if res {
		t.Fail()
	}
}

func TestAllPtrTrue(t *testing.T) {
	res := AllPtr(func(i *int) bool { return *i > 5 }, []int{6, 7, 8, 9})
	if !res {
		t.Fail()
	}
}

func TestAllPtrFalse(t *testing.T) {
	res := AllPtr(func(i *int) bool { return *i > 5 }, []int{6, 7, 8, 9, 1})
	if res {
		t.Fail()
	}
}
