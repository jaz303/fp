package fp

func Zip[A, B any](as []A, bs []B) []Pair[A, B] {
	out := make([]Pair[A, B], len(as))
	for i := range as {
		out[i] = Pair[A, B]{as[i], bs[i]}
	}
	return out
}
