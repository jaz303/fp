package fp

import "testing"

func TestFilter(t *testing.T) {
	out := Filter(func(i int) bool {
		return i > 2
	}, []int{1, 2, 3, 4})
	if !Eq(out, []int{3, 4}) {
		t.Fail()
	}
}

func TestFilterPtr(t *testing.T) {
	out := FilterPtr(func(i *int) bool {
		return *i > 2
	}, []int{1, 2, 3, 4})
	if !Eq(out, []int{3, 4}) {
		t.Fail()
	}
}

func TestFilterIndex(t *testing.T) {
	out := FilterIndex(func(i int, ix int) bool {
		return i > 2 && ix > 2
	}, []int{1, 2, 3, 4})
	if !Eq(out, []int{4}) {
		t.Fail()
	}
}

func TestFilterIndexPtr(t *testing.T) {
	out := FilterIndexPtr(func(i *int, ix int) bool {
		return *i > 2 && ix > 2
	}, []int{1, 2, 3, 4})
	if !Eq(out, []int{4}) {
		t.Fail()
	}
}
