package fp

func All[I any](fn func(I) bool, in []I) bool {
	for _, v := range in {
		if !fn(v) {
			return false
		}
	}
	return true
}

func AllPtr[I any](fn func(*I) bool, in []I) bool {
	for i := range in {
		if !fn(&in[i]) {
			return false
		}
	}
	return true
}
