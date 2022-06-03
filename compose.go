package fp

func Compose[T1, T2, T3 any](fn1 func(T1) T2, fn2 func(T2) T3) func(T1) T3 {
	return func(v T1) T3 {
		return fn2(fn1(v))
	}
}
