package fp

import "testing"

func TestUnzip(t *testing.T) {
	as, bs := Unzip([]Pair[int, string]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	})

	if !Eq(as, []int{1, 2, 3}) {
		t.Fail()
	}

	if !Eq(bs, []string{"a", "b", "c"}) {
		t.Fail()
	}
}
