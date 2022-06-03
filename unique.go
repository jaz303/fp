package fp

func Unique[I comparable](in []I) []I {
	seen := make(map[I]struct{})
	for _, v := range in {
		seen[v] = struct{}{}
	}
	return Keys(seen)
}
