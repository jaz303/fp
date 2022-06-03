package fp

func Count[I any](fn func(I) bool, in []I) int {
	c := 0
	for _, v := range in {
		if fn(v) {
			c++
		}
	}
	return c
}

func CountPtr[I any](fn func(*I) bool, in []I) int {
	c := 0
	for i := range in {
		if fn(&in[i]) {
			c++
		}
	}
	return c
}

func CountIndex[I any](fn func(I, int) bool, in []I) int {
	c := 0
	for i, v := range in {
		if fn(v, i) {
			c++
		}
	}
	return c
}

func CountIndexPtr[I any](fn func(*I, int) bool, in []I) int {
	c := 0
	for i := range in {
		if fn(&in[i], i) {
			c++
		}
	}
	return c
}
