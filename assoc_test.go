package fp

import "testing"

func TestAssoc(t *testing.T) {
	res := Assoc([]int{1, 2, 3}, []string{"a", "b", "c"})

	if len(res) != 3 || res[1] != "a" || res[2] != "b" || res[3] != "c" {
		t.Fail()
	}
}
