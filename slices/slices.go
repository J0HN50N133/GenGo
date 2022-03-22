package slices

func Map[A, B any](f func(A) B, l []A) []B {
	res := make([]B, len(l))
	for i, e := range l {
		res[i] = f(e)
	}
	return res
}

func ForEach[A any](f func(A), l []A) {
	for _, e := range l {
		f(e)
	}
}

func Foldl[A, B any](f func(B, A) B, b B, l []A) B {
	res := b
	for _, e := range l {
		res = f(res, e)
	}
	return res
}

func Foldr[A, B any](f func(A, B) B, b B, l []A) B {
	res := b
	for i := len(l) - 1; i >= 0; i-- {
		res = f(l[i], res)
	}
	return res
}

func Filter[A any](f func(A) bool, l []A) []A {
	res := make([]A, 0, len(l))
	for _, e := range l {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}

func Any[A any](f func(A) bool, l []A) bool {
	for _, e := range l {
		if f(e) {
			return true
		}
	}
	return false
}

func Concat[A any](l [][]A) []A {
	return Foldl(func(a []A, b []A) []A { return append(a, b...) }, []A{}, l)
}

func ConcatMap[A, B any](f func(A) []B, l []A) []B {
	return Concat(Map(f, l))
}
