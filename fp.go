package fp

func Assoc[K comparable, V any](keys []K, values []V) map[K]V {
	out := make(map[K]V, len(keys))
	for i, k := range keys {
		out[k] = values[i]
	}
	return out
}

func Compose[T1, T2, T3 any](fn1 func(T1) T2, fn2 func(T2) T3) func(T1) T3 {
	return func(v T1) T3 {
		return fn2(fn1(v))
	}
}

func Foreach[I any](fn func(I), in []I) {
	for _, v := range in {
		fn(v)
	}
}

func ForeachIndex[I any](fn func(I, int), in []I) {
	for i, v := range in {
		fn(v, i)
	}
}

func Map[I, V any](fn func(I) V, in []I) []V {
	out := make([]V, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func MapIndex[I, V any](fn func(I, int) V, in []I) []V {
	out := make([]V, len(in))
	for i, v := range in {
		out[i] = fn(v, i)
	}
	return out
}

func Reduce[I, V any](fn func(acc V, val I) V, in []I, initial V) V {
	for _, v := range in {
		initial = fn(initial, v)
	}
	return initial
}

func ReduceIndex[I, V any](fn func(acc V, val I, index int) V, in []I, initial V) V {
	for i, v := range in {
		initial = fn(initial, v, i)
	}
	return initial
}

func Repeat[I any](v I, n int) []I {
	out := make([]I, n)
	for i := range out {
		out[i] = v
	}
	return out
}

func Keys[K comparable, V any](in map[K]V) []K {
	out := make([]K, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

func Values[K comparable, V any](in map[K]V) []V {
	out := make([]V, len(in))
	for k := range in {
		out = append(out, in[k])
	}
	return out
}

func Unique[I comparable](in []I) []I {
	seen := make(map[I]struct{})
	for _, v := range in {
		seen[v] = struct{}{}
	}
	return Keys(seen)
}

func All[I any](fn func(I) bool, in []I) bool {
	for _, v := range in {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Any[I any](fn func(I) bool, in []I) bool {
	for _, v := range in {
		if fn(v) {
			return true
		}
	}
	return false
}

func None[I any](fn func(I) bool, in []I) bool {
	return !Any(fn, in)
}

func First[I any](fn func(I) bool, in []I) (int, I) {
	for i, v := range in {
		if fn(v) {
			return i, v
		}
	}
	var zero I
	return -1, zero
}

func Count[I any](fn func(I) bool, in []I) int {
	c := 0
	for _, v := range in {
		if fn(v) {
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

func Filter[I any](fn func(I) bool, in []I) []I {
	out := make([]I, 0)
	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func FilterIndex[I any](fn func(I, int) bool, in []I) []I {
	out := make([]I, 0)
	for i, v := range in {
		if fn(v, i) {
			out = append(out, v)
		}
	}
	return out
}

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
