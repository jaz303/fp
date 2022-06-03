package fp

import (
	"fmt"
	"testing"
)

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
