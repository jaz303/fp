package fp

func Reduce[I, V any](fn func(acc V, val I) V, in []I, initial V) V {
	for _, v := range in {
		initial = fn(initial, v)
	}
	return initial
}

func ReducePtr[I, V any](fn func(acc V, val *I) V, in []I, initial V) V {
	for i := range in {
		initial = fn(initial, &in[i])
	}
	return initial
}

func ReduceIndex[I, V any](fn func(acc V, val I, index int) V, in []I, initial V) V {
	for i, v := range in {
		initial = fn(initial, v, i)
	}
	return initial
}

func ReduceIndexPtr[I, V any](fn func(acc V, val *I, index int) V, in []I, initial V) V {
	for i := range in {
		initial = fn(initial, &in[i], i)
	}
	return initial
}
