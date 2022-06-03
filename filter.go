package fp

func Filter[I any](fn func(I) bool, in []I) []I {
	out := make([]I, 0)
	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func FilterPtr[I any](fn func(*I) bool, in []I) []I {
	out := make([]I, 0)
	for i := range in {
		if fn(&in[i]) {
			out = append(out, in[i])
		}
	}
	return out
}

func FilterIndex[I any](fn func(I, int) bool, in []I) []I {
	out := make([]I, 0)
	for i, v := range in {
		if fn(v, i) {
			out = append(out, v)
		}
	}
	return out
}

func FilterIndexPtr[I any](fn func(*I, int) bool, in []I) []I {
	out := make([]I, 0)
	for i := range in {
		if fn(&in[i], i) {
			out = append(out, in[i])
		}
	}
	return out
}
