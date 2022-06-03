package fp

import "testing"

func TestRepeat(t *testing.T) {
	res := Repeat("hi", 3)

	if len(res) != 3 || res[0] != "hi" || res[1] != "hi" || res[2] != "hi" {
		t.Fail()
	}
}
