package fp

import "testing"

func TestPartition(t *testing.T) {
	odds, evens := Partition(func(i int) bool {
		return (i % 2) == 1
	}, []int{1, 2, 3, 4, 5, 6})

	if !Eq(odds, []int{1, 3, 5}) {
		t.Fail()
	}

	if !Eq(evens, []int{2, 4, 6}) {
		t.Fail()
	}
}

func TestPartitionPtr(t *testing.T) {
	odds, evens := PartitionPtr(func(i *int) bool {
		return (*i % 2) == 1
	}, []int{1, 2, 3, 4, 5, 6})

	if !Eq(odds, []int{1, 3, 5}) {
		t.Fail()
	}

	if !Eq(evens, []int{2, 4, 6}) {
		t.Fail()
	}
}

func TestPartitionIndex(t *testing.T) {
	firstTwo, rest := PartitionIndex(func(i int, ix int) bool {
		return ix < 2
	}, []int{1, 2, 3, 4, 5, 6})

	if !Eq(firstTwo, []int{1, 2}) {
		t.Fail()
	}

	if !Eq(rest, []int{3, 4, 5, 6}) {
		t.Fail()
	}
}

func TestPartitionIndexPtr(t *testing.T) {
	firstTwo, rest := PartitionIndexPtr(func(i *int, ix int) bool {
		return ix < 2
	}, []int{1, 2, 3, 4, 5, 6})

	if !Eq(firstTwo, []int{1, 2}) {
		t.Fail()
	}

	if !Eq(rest, []int{3, 4, 5, 6}) {
		t.Fail()
	}
}
