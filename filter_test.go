package fp

import (
	"testing"

	"github.com/jaz303/fp/pred"
)

func TestFilter(t *testing.T) {
	out := Filter(pred.Gt(2), []int{1, 2, 3, 4})
	if !Eq(out, []int{3, 4}) {
		t.Fail()
	}
}

func TestFilterPtr(t *testing.T) {
	out := FilterPtr(pred.GtPtr(2), []int{1, 2, 3, 4})
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
