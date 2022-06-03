package fp

import (
	"fmt"
	"testing"
)

func TestAssoc(t *testing.T) {
	res := Assoc([]int{1, 2, 3}, []string{"a", "b", "c"})

	if len(res) != 3 || res[1] != "a" || res[2] != "b" || res[3] != "c" {
		t.Fail()
	}
}

func TestCompose(t *testing.T) {
	fn := Compose(
		func(i int) string { return fmt.Sprintf("%d", i) },
		func(s string) bool { return len(s) > 4 },
	)

	res := fn(10000)

	if !res {
		t.Fail()
	}
}

func TestForeach(t *testing.T) {
	var res []int
	Foreach(func(i int) {
		res = append(res, i)
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

func TestMap(t *testing.T) {
	res := Map(func(i int) string {
		return fmt.Sprintf("%d", i*2)
	}, []int{1, 2, 3})

	if len(res) != 3 || res[0] != "2" || res[1] != "4" || res[2] != "6" {
		t.Fail()
	}
}

func TestMapIndex(t *testing.T) {
	res := MapIndex(func(i int, ix int) int {
		return (2 * i) + ix
	}, []int{1, 2, 3})

	if len(res) != 3 || res[0] != 2 || res[1] != 5 || res[2] != 8 {
		t.Fail()
	}
}

func TestReduce(t *testing.T) {
	sum := Reduce(func(sum int, i int) int {
		return sum + i
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

func TestRepeat(t *testing.T) {
	res := Repeat("hi", 3)

	if len(res) != 3 || res[0] != "hi" || res[1] != "hi" || res[2] != "hi" {
		t.Fail()
	}
}

func TestAllTrue(t *testing.T) {
	res := All(func(i int) bool { return i > 5 }, []int{6, 7, 8, 9})
	if !res {
		t.Fail()
	}
}

func TestAllFalse(t *testing.T) {
	res := All(func(i int) bool { return i > 5 }, []int{6, 7, 8, 9, 1})
	if res {
		t.Fail()
	}
}

func TestAnyTrue(t *testing.T) {
	res := Any(func(i int) bool { return i > 5 }, []int{0, -1, 5, 10})
	if !res {
		t.Fail()
	}
}

func TestAnyFalse(t *testing.T) {
	res := Any(func(i int) bool { return i > 5 }, []int{-1, -2, -3, 4})
	if res {
		t.Fail()
	}
}

func TestNoneTrue(t *testing.T) {
	res := None(func(i int) bool { return i > 5 }, []int{0, -1, 2, 3})
	if !res {
		t.Fail()
	}
}

func TestNoneFalse(t *testing.T) {
	res := None(func(i int) bool { return i > 5 }, []int{-1, -2, 10, 4})
	if res {
		t.Fail()
	}
}
