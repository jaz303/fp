package fp

import "testing"

func TestValues(t *testing.T) {
	values := Values(map[string]int{"a": 1, "b": 2})

	if len(values) != 2 {
		t.Fail()
	}

	got1 := false
	got2 := false
	for _, k := range values {
		if k == 1 {
			got1 = true
		} else if k == 2 {
			got2 = true
		}
	}

	if !got1 || !got2 {
		t.Fail()
	}
}
