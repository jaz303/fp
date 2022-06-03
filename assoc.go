package fp

func Assoc[K comparable, V any](keys []K, values []V) map[K]V {
	out := make(map[K]V, len(keys))
	for i, k := range keys {
		out[k] = values[i]
	}
	return out
}
