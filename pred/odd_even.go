package pred

import "golang.org/x/exp/constraints"

func Odd[T constraints.Integer](v T) bool {
	return v%2 == 1
}

func Even[T constraints.Integer](v T) bool {
	return v%2 == 0
}
