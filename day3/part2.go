package day3

import (
	"fmt"
	"math/bits"
	"strconv"
)

func (r *DiagnosticReport) AddNumber(value uint64) {
	r.numbers = append(r.numbers, value)

	if bits.Len64(value) > r.length {
		r.length = bits.Len64(value)
	}
}

func (r *DiagnosticReport) Print() {
	for _, number := range r.numbers {
		FastPrintBinary(number)
	}
	fmt.Println()
}

func (r *DiagnosticReport) CalculateOxygenGeneratorRating(mask uint64) uint64 {
	if len(r.numbers) <= 1 {
		return r.numbers[0]
	}

	var zeros, ones, candidates DiagnosticReport

	for _, number := range r.numbers {
		if mask&number == 0 {
			zeros.AddNumber(number)
		} else {
			ones.AddNumber(number)
		}
	}

	if len(zeros.numbers) == 0 {
		candidates = ones
	} else if len(ones.numbers) == 0 {
		candidates = zeros
	} else if len(ones.numbers) > 0 && len(ones.numbers) >= len(zeros.numbers) {
		candidates = ones
	} else {
		candidates = zeros
	}

	return candidates.CalculateOxygenGeneratorRating(mask >> 1)
}

func FastPrintBinary(num uint64) {
	fmt.Println(strconv.FormatUint(num, 2))
}

func (r *DiagnosticReport) CalculateCO2ScrubberRating(mask uint64) uint64 {
	if len(r.numbers) <= 1 {
		return r.numbers[0]
	}

	var zeros, ones, candidates DiagnosticReport

	for _, number := range r.numbers {
		if mask&number == 0 {
			zeros.AddNumber(number)
		} else {
			ones.AddNumber(number)
		}
	}

	if len(zeros.numbers) == 0 {
		candidates = ones
	} else if len(ones.numbers) == 0 {
		candidates = zeros
	} else if len(zeros.numbers) <= len(ones.numbers) {
		candidates = zeros
	} else {
		candidates = ones
	}

	return candidates.CalculateCO2ScrubberRating(mask >> 1)
}

func (r *DiagnosticReport) CalculateLifeSupportRating() uint64 {
	co2ScrubberRating := r.CalculateCO2ScrubberRating(1 << r.length)
	oxygenGeneratorRating := r.CalculateOxygenGeneratorRating(1 << r.length)
	return co2ScrubberRating * oxygenGeneratorRating
}
