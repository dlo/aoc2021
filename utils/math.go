package utils

func SumSlice(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
