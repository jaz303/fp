package fp

import "testing"

func TestReduce(t *testing.T) {
	sum := Reduce(func(sum int, i int) int {
		return sum + i
	}, []int{1, 2, 3}, 0)

	if sum != 6 {
		t.Fail()
	}
}

func TestReducePtr(t *testing.T) {
	sum := ReducePtr(func(sum int, i *int) int {
		return sum + *i
	}, []int{1, 2, 3}, 0)

	if sum != 6 {
		t.Fail()
	}
}

func TestReduceIndex(t *testing.T) {
	sum := ReduceIndex(func(sum int, i int, ix int) int {
		return sum + i + ix
	}, []int{1, 2, 3}, 0)

	if sum != 9 {
		t.Fail()
	}
}

func TestReduceIndexPtr(t *testing.T) {
	sum := ReduceIndexPtr(func(sum int, i *int, ix int) int {
		return sum + *i + ix
	}, []int{1, 2, 3}, 0)

	if sum != 9 {
		t.Fail()
	}
}
