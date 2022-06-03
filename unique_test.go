package fp

import "testing"

func TestUnique(t *testing.T) {
	u := Unique([]int{1, 2, 1, 2, 1, 1, 2, 2})

	if len(u) != 2 {
		t.Fail()
	}

	got1 := false
	got2 := false
	for _, v := range u {
		if v == 1 {
			got1 = true
		}
		if v == 2 {
			got2 = true
		}
	}

	if !got1 || !got2 {
		t.Fail()
	}
}
