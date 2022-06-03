package fp

func Values[K comparable, V any](in map[K]V) []V {
	out := make([]V, 0, len(in))
	for k := range in {
		out = append(out, in[k])
	}
	return out
}
