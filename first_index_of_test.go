package fp

import (
	"testing"

	"github.com/jaz303/fp/pred"
)

func TestFirstIndexOfPresent(t *testing.T) {
	ix := FirstIndexOf(pred.Gt(10), []int{0, 5, 10, 15})
	if ix != 3 {
		t.Fail()
	}
}

func TestFirstIndexOfMissing(t *testing.T) {
	ix := FirstIndexOf(pred.Gt(100), []int{0, 5, 10, 15})
	if ix != -1 {
		t.Fail()
	}
}

func TestFirstIndexOfPtrPresent(t *testing.T) {
	ix := FirstIndexOfPtr(pred.GtPtr(10), []int{0, 5, 10, 15})
	if ix != 3 {
		t.Fail()
	}
}

func TestFirstIndexOfPtrMissing(t *testing.T) {
	ix := FirstIndexOfPtr(pred.GtPtr(100), []int{0, 5, 10, 15})
	if ix != -1 {
		t.Fail()
	}
}
