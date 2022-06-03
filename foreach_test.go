package fp

import "testing"

func TestForeach(t *testing.T) {
	var res []int
	Foreach(func(i int) {
		res = append(res, i)
	}, []int{1, 2, 3})

	if len(res) != 3 || res[0] != 1 || res[1] != 2 || res[2] != 3 {
		t.Fail()
	}
}

func TestForeachPtr(t *testing.T) {
	var res []int
	ForeachPtr(func(i *int) {
		res = append(res, *i)
	}, []int{1, 2, 3})

	if len(res) != 3 || res[0] != 1 || res[1] != 2 || res[2] != 3 {
		t.Fail()
	}
}

func TestForeachIndex(t *testing.T) {
	var res []int
	var ind []int
	ForeachIndex(func(i int, ix int) {
		res = append(res, i)
		ind = append(ind, ix)
	}, []int{1, 2, 3})

	if len(res) != 3 || res[0] != 1 || res[1] != 2 || res[2] != 3 {
		t.Fail()
	}

	if len(ind) != 3 || ind[0] != 0 || ind[1] != 1 || ind[2] != 2 {
		t.Fail()
	}
}

func TestForeachIndexPtr(t *testing.T) {
	var res []int
	var ind []int
	ForeachIndexPtr(func(i *int, ix int) {
		res = append(res, *i)
		ind = append(ind, ix)
	}, []int{1, 2, 3})

	if len(res) != 3 || res[0] != 1 || res[1] != 2 || res[2] != 3 {
		t.Fail()
	}

	if len(ind) != 3 || ind[0] != 0 || ind[1] != 1 || ind[2] != 2 {
		t.Fail()
	}
}
