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

func (r *DiagnosticReport) GroupNumbersBasedOnMask(mask uint64) (*DiagnosticReport, *DiagnosticReport) {
	var zeros, ones DiagnosticReport
	for _, number := range r.numbers {
		if mask&number == 0 {
			zeros.AddNumber(number)
		} else {
			ones.AddNumber(number)
		}
	}
	return &zeros, &ones
}

func (r *DiagnosticReport) CalculateRating(mask uint64, criteria RatingCriteria) uint64 {
	if len(r.numbers) <= 1 {
		return r.numbers[0]
	}

	zeros, ones := r.GroupNumbersBasedOnMask(mask)
	var candidates DiagnosticReport

	if len(ones.numbers) == 0 || len(zeros.numbers) == 0 {
		candidates = *r
	} else {
		switch criteria {
		case OxygenGeneratorRatingCriteria:
			if len(ones.numbers) >= len(zeros.numbers) {
				candidates = *ones
			} else {
				candidates = *zeros
			}
		case CO2RatingCriteria:
			if len(zeros.numbers) <= len(ones.numbers) {
				candidates = *zeros
			} else {
				candidates = *ones
			}
		}
	}

	return candidates.CalculateRating(mask>>1, criteria)
}

func (r *DiagnosticReport) CalculateOxygenGeneratorRating(mask uint64) uint64 {
	return r.CalculateRating(mask, OxygenGeneratorRatingCriteria)
}

func (r *DiagnosticReport) CalculateCO2ScrubberRating(mask uint64) uint64 {
	return r.CalculateRating(mask, CO2RatingCriteria)
}

func (r *DiagnosticReport) CalculateLifeSupportRating() uint64 {
	co2ScrubberRating := r.CalculateCO2ScrubberRating(1 << r.length)
	oxygenGeneratorRating := r.CalculateOxygenGeneratorRating(1 << r.length)
	return co2ScrubberRating * oxygenGeneratorRating
}
