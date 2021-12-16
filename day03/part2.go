package day3

import (
	"fmt"
	"math/bits"
	"strconv"
)

type RatingCriteria int

const (
	OxygenGeneratorRatingCriteria = iota
	CO2RatingCriteria
)

func FastPrintBinary(num uint64) {
	fmt.Println(strconv.FormatUint(num, 2))
}

func (r *DiagnosticReport) AddNumber(value uint64) {
	r.Numbers = append(r.Numbers, value)

	if bits.Len64(value) > r.Length {
		r.Length = bits.Len64(value)
	}
}

func (r *DiagnosticReport) Print() {
	for _, number := range r.Numbers {
		FastPrintBinary(number)
	}
	fmt.Println()
}

func (r *DiagnosticReport) GroupNumbersBasedOnMask(mask uint64) (*DiagnosticReport, *DiagnosticReport) {
	var zeros, ones DiagnosticReport
	for _, number := range r.Numbers {
		if mask&number == 0 {
			zeros.AddNumber(number)
		} else {
			ones.AddNumber(number)
		}
	}
	return &zeros, &ones
}

func (r *DiagnosticReport) CalculateRating(mask uint64, criteria RatingCriteria) uint64 {
	if len(r.Numbers) <= 1 {
		return r.Numbers[0]
	}

	zeros, ones := r.GroupNumbersBasedOnMask(mask)

	if len(ones.Numbers) == 0 || len(zeros.Numbers) == 0 {
		return r.CalculateRating(mask>>1, criteria)
	} else {
		if criteria == OxygenGeneratorRatingCriteria && len(zeros.Numbers) > len(ones.Numbers) ||
			criteria == CO2RatingCriteria && len(zeros.Numbers) <= len(ones.Numbers) {
			return zeros.CalculateRating(mask>>1, criteria)
		} else {
			return ones.CalculateRating(mask>>1, criteria)
		}
	}
}

func (r *DiagnosticReport) CalculateOxygenGeneratorRating(mask uint64) uint64 {
	return r.CalculateRating(mask, OxygenGeneratorRatingCriteria)
}

func (r *DiagnosticReport) CalculateCO2ScrubberRating(mask uint64) uint64 {
	return r.CalculateRating(mask, CO2RatingCriteria)
}

func (r *DiagnosticReport) CalculateLifeSupportRating() uint64 {
	co2ScrubberRating := r.CalculateCO2ScrubberRating(1 << r.Length)
	oxygenGeneratorRating := r.CalculateOxygenGeneratorRating(1 << r.Length)
	return co2ScrubberRating * oxygenGeneratorRating
}
