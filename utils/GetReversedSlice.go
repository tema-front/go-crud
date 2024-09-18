package utils

func GetReversedSlice[T any](s []T) []T {
	sLen := len(s)
	rs := make([]T, sLen)

	for i, v := range s {
		ri := sLen - i - 1
		rs[ri] = v
	}

	return rs
}

//