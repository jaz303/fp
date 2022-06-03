package fp

import (
	"testing"

	"github.com/jaz303/fp/pred"
)

func TestAllTrue(t *testing.T) {
	res := All(pred.Gt(5), []int{6, 7, 8, 9})
	if !res {
		t.Fail()
	}
}

func TestAllFalse(t *testing.T) {
	res := All(pred.Gt(5), []int{6, 7, 8, 9, 1})
	if res {
		t.Fail()
	}
}

func TestAllPtrTrue(t *testing.T) {
	res := AllPtr(pred.GtPtr(5), []int{6, 7, 8, 9})
	if !res {
		t.Fail()
	}
}

func TestAllPtrFalse(t *testing.T) {
	res := AllPtr(pred.GtPtr(5), []int{6, 7, 8, 9, 1})
	if res {
		t.Fail()
	}
}
