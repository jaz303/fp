package fp

import (
	"testing"

	"github.com/jaz303/fp/pred"
)

func TestAnyTrue(t *testing.T) {
	res := Any(pred.Gt(5), []int{0, -1, 5, 10})
	if !res {
		t.Fail()
	}
}

func TestAnyFalse(t *testing.T) {
	res := Any(pred.Gt(5), []int{-1, -2, -3, 4})
	if res {
		t.Fail()
	}
}

func TestAnyPtrTrue(t *testing.T) {
	res := AnyPtr(pred.GtPtr(5), []int{0, -1, 5, 10})
	if !res {
		t.Fail()
	}
}

func TestAnyPtrFalse(t *testing.T) {
	res := AnyPtr(pred.GtPtr(5), []int{-1, -2, -3, 4})
	if res {
		t.Fail()
	}
}
