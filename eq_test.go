package fp

import "testing"

func TestEqTrue(t *testing.T) {
	if !Eq([]int{}, []int{}) {
		t.Fail()
	}

	if !Eq([]string{"a", "b", "c"}, []string{"a", "b", "c"}) {
		t.Fail()
	}
}

func TestEqFalse(t *testing.T) {
	if Eq([]int{1, 2}, []int{1, 2, 3}) {
		t.Fail()
	}

	if Eq([]string{"a", "d", "c"}, []string{"a", "b", "c"}) {
		t.Fail()
	}
}
