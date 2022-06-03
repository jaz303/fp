package fp

func FirstIndexOf[I any](fn func(I) bool, in []I) int {
	for i, v := range in {
		if fn(v) {
			return i
		}
	}
	return -1
}

func FirstIndexOfPtr[I any](fn func(*I) bool, in []I) int {
	for i := range in {
		if fn(&in[i]) {
			return i
		}
	}
	return -1
}
