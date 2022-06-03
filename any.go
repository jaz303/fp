package fp

func Any[I any](fn func(I) bool, in []I) bool {
	for _, v := range in {
		if fn(v) {
			return true
		}
	}
	return false
}

func AnyPtr[I any](fn func(*I) bool, in []I) bool {
	for i := range in {
		if fn(&in[i]) {
			return true
		}
	}
	return false
}
