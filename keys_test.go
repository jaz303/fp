package fp

import "testing"

func TestKeys(t *testing.T) {
	keys := Keys(map[string]int{"a": 1, "b": 2})

	if len(keys) != 2 {
		t.Fail()
	}

	gotA := false
	gotB := false
	for _, k := range keys {
		if k == "a" {
			gotA = true
		} else if k == "b" {
			gotB = true
		}
	}

	if !gotA || !gotB {
		t.Fail()
	}
}
