package fp

import "testing"

func TestFirstPresent(t *testing.T) {
	v, ix := First(func(i int) bool {
		return i > 10
	}, []int{5, 10, 15, 20})
	if v != 15 || ix != 2 {
		t.Fail()
	}
}

func TestFirstMissing(t *testing.T) {
	v, ix := First(func(i int) bool {
		return i > 100
	}, []int{5, 10, 15, 20})
	if v != 0 || ix != -1 {
		t.Fail()
	}
}

func TestFirstPtrPresent(t *testing.T) {
	v, ix := FirstPtr(func(i *int) bool {
		return *i > 10
	}, []int{5, 10, 15, 20})
	if *v != 15 || ix != 2 {
		t.Fail()
	}
}

func TestFirstPtrMissing(t *testing.T) {
	v, ix := FirstPtr(func(i *int) bool {
		return *i > 100
	}, []int{5, 10, 15, 20})
	if v != nil || ix != -1 {
		t.Fail()
	}
}

func TestFirstIndexPresent(t *testing.T) {
	v, ix := FirstIndex(func(i int, ix int) bool {
		return i > 10 && ix > 2
	}, []int{5, 10, 15, 20})
	if v != 20 || ix != 3 {
		t.Fail()
	}
}

func TestFirstIndexMissing(t *testing.T) {
	v, ix := FirstIndex(func(i int, ix int) bool {
		return i > 0 && ix > 5
	}, []int{5, 10, 15, 20})
	if v != 0 || ix != -1 {
		t.Fail()
	}
}

func TestFirstIndexPtrPresent(t *testing.T) {
	v, ix := FirstIndexPtr(func(i *int, ix int) bool {
		return *i > 10 && ix > 2
	}, []int{5, 10, 15, 20})
	if *v != 20 || ix != 3 {
		t.Fail()
	}
}

func TestFirstIndexPtrMissing(t *testing.T) {
	v, ix := FirstIndexPtr(func(i *int, ix int) bool {
		return *i > 0 && ix > 5
	}, []int{5, 10, 15, 20})
	if v != nil || ix != -1 {
		t.Fail()
	}
}
