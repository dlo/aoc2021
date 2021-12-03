package day3

import (
	"math/bits"
)

type DiagnosticReport struct {
	numbers []uint64
	length  int
}

func (r *DiagnosticReport) CalculatePowerConsumption() uint64 {
	// `sums` stores the most common bits in reverse order
	sums := make([]float64, r.length)
	for _, number := range r.numbers {
		var mask uint64 = 1
		for i := 0; i < len(sums); i++ {
			value := float64(bits.OnesCount64(mask&number)) - 0.5
			sums[i] += value
			mask <<= 1
		}
	}

	var gammaRate uint64 = 0
	var epsilonRate uint64 = 0
	var mask uint64 = 1
	for _, value := range sums {
		if value > 0 {
			gammaRate += mask
		} else {
			epsilonRate += mask
		}
		mask <<= 1
	}

	return gammaRate * epsilonRate
}
