package fp

func First[I any](fn func(I) bool, in []I) (I, int) {
	for i, v := range in {
		if fn(v) {
			return v, i
		}
	}
	var zero I
	return zero, -1
}

func FirstPtr[I any](fn func(*I) bool, in []I) (*I, int) {
	for i := range in {
		if fn(&in[i]) {
			return &in[i], i
		}
	}
	return nil, -1
}

func FirstIndex[I any](fn func(I, int) bool, in []I) (I, int) {
	for i, v := range in {
		if fn(v, i) {
			return v, i
		}
	}
	var zero I
	return zero, -1
}

func FirstIndexPtr[I any](fn func(*I, int) bool, in []I) (*I, int) {
	for i := range in {
		if fn(&in[i], i) {
			return &in[i], i
		}
	}
	return nil, -1
}
