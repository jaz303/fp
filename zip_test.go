package fp

import "testing"

func TestZip(t *testing.T) {
	z := Zip([]int{1, 2, 3}, []string{"a", "b", "c"})
	if !Eq(z, []Pair[int, string]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}) {
		t.Fail()
	}
}
