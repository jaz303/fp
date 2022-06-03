package fp

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	res := Map(func(i int) string {
		return fmt.Sprintf("%d", i*2)
	}, []int{1, 2, 3})

	if len(res) != 3 || res[0] != "2" || res[1] != "4" || res[2] != "6" {
		t.Fail()
	}
}

func TestMapPtr(t *testing.T) {
	res := MapPtr(func(i *int) string {
		return fmt.Sprintf("%d", (2 * *i))
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

func TestMapIndexPtr(t *testing.T) {
	res := MapIndexPtr(func(i *int, ix int) int {
		return (2 * *i) + ix
	}, []int{1, 2, 3})

	if len(res) != 3 || res[0] != 2 || res[1] != 5 || res[2] != 8 {
		t.Fail()
	}
}
