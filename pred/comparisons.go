package pred

import "golang.org/x/exp/constraints"

func Zero[T comparable](v T) bool {
	var zero T
	return v == zero
}

func Gt[T constraints.Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x > v
	}
}

func GtPtr[T constraints.Ordered](v T) func(*T) bool {
	return func(x *T) bool {
		return *x > v
	}
}

func GtEq[T constraints.Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x >= v
	}
}

func GtEqPtr[T constraints.Ordered](v T) func(*T) bool {
	return func(x *T) bool {
		return *x >= v
	}
}

func Lt[T constraints.Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x < v
	}
}

func LtPtr[T constraints.Ordered](v T) func(*T) bool {
	return func(x *T) bool {
		return *x < v
	}
}

func LtEq[T constraints.Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x <= v
	}
}

func LtEqPtr[T constraints.Ordered](v T) func(*T) bool {
	return func(x *T) bool {
		return *x <= v
	}
}

func Eq[T comparable](v T) func(T) bool {
	return func(x T) bool {
		return x == v
	}
}

func EqPtr[T comparable](v T) func(*T) bool {
	return func(x *T) bool {
		return *x == v
	}
}

func NotEq[T comparable](v T) func(T) bool {
	return func(x T) bool {
		return x != v
	}
}

func NotEqPtr[T comparable](v T) func(*T) bool {
	return func(x *T) bool {
		return *x != v
	}
}
