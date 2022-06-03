package fp

import (
	"testing"

	"github.com/jaz303/fp/pred"
)

func TestCount(t *testing.T) {
	c := Count(pred.Gt(10), []int{1, 2, 11, 12})

	if c != 2 {
		t.Fail()
	}
}

func TestCountPtr(t *testing.T) {
	c := CountPtr(pred.GtPtr(10), []int{1, 2, 11, 12})

	if c != 2 {
		t.Fail()
	}
}

func TestCountIndex(t *testing.T) {
	c := CountIndex(func(i int, ix int) bool {
		return i > 10 || ix > 0
	}, []int{1, 2, 11, 12})

	if c != 3 {
		t.Fail()
	}
}

func TestCountIndexPtr(t *testing.T) {
	c := CountIndexPtr(func(i *int, ix int) bool {
		return *i > 10 || ix > 0
	}, []int{1, 2, 11, 12})

	if c != 3 {
		t.Fail()
	}
}
