package fp

func None[I any](fn func(I) bool, in []I) bool {
	return !Any(fn, in)
}

func NonePtr[I any](fn func(*I) bool, in []I) bool {
	return !AnyPtr(fn, in)
}
