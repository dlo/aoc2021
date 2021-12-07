package utils

func SliceSum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func SliceAverage(values []int) float64 {
	sum := SliceSum(values)
	return float64(sum) / float64(len(values))
}

func Abs(value int) int {
	if value > 0 {
		return value
	}

	return -1 * value
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
