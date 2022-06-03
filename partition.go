package fp

func Partition[I any](fn func(I) bool, in []I) ([]I, []I) {
	yes, no := make([]I, 0), make([]I, 0)
	for _, v := range in {
		if fn(v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return yes, no
}

func PartitionPtr[I any](fn func(*I) bool, in []I) ([]I, []I) {
	yes, no := make([]I, 0), make([]I, 0)
	for i := range in {
		if fn(&in[i]) {
			yes = append(yes, in[i])
		} else {
			no = append(no, in[i])
		}
	}
	return yes, no
}

func PartitionIndex[I any](fn func(I, int) bool, in []I) ([]I, []I) {
	yes, no := make([]I, 0), make([]I, 0)
	for i, v := range in {
		if fn(v, i) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return yes, no
}

func PartitionIndexPtr[I any](fn func(*I, int) bool, in []I) ([]I, []I) {
	yes, no := make([]I, 0), make([]I, 0)
	for i := range in {
		if fn(&in[i], i) {
			yes = append(yes, in[i])
		} else {
			no = append(no, in[i])
		}
	}
	return yes, no
}
