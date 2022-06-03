package fp

func Unzip[A, B any](list []Pair[A, B]) ([]A, []B) {
	as := make([]A, len(list))
	bs := make([]B, len(list))

	for i := range list {
		as[i] = list[i].A
		bs[i] = list[i].B
	}

	return as, bs
}
