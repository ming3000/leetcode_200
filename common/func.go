package common

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(data int) int {
	if data < 0 {
		return -1 * data
	}
	return data
}
