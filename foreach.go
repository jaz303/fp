package fp

func Foreach[I any](fn func(I), in []I) {
	for _, v := range in {
		fn(v)
	}
}

func ForeachPtr[I any](fn func(*I), in []I) {
	for i := range in {
		fn(&in[i])
	}
}

func ForeachIndex[I any](fn func(I, int), in []I) {
	for i, v := range in {
		fn(v, i)
	}
}

func ForeachIndexPtr[I any](fn func(*I, int), in []I) {
	for i := range in {
		fn(&in[i], i)
	}
}
