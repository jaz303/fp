package fp

func Map[I, V any](fn func(I) V, in []I) []V {
	out := make([]V, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func MapPtr[I, V any](fn func(*I) V, in []I) []V {
	out := make([]V, len(in))
	for i := range in {
		out[i] = fn(&in[i])
	}
	return out
}

func MapIndex[I, V any](fn func(I, int) V, in []I) []V {
	out := make([]V, len(in))
	for i, v := range in {
		out[i] = fn(v, i)
	}
	return out
}

func MapIndexPtr[I, V any](fn func(*I, int) V, in []I) []V {
	out := make([]V, len(in))
	for i := range in {
		out[i] = fn(&in[i], i)
	}
	return out
}
