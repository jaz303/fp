package fp

func Repeat[I any](v I, n int) []I {
	out := make([]I, n)
	for i := range out {
		out[i] = v
	}
	return out
}
