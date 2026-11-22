package main

import (
	"errors"
	"fmt"
)

const (
	minimumTemperature = 15
	maximumTemperature = 30
)

var errInvalidComparison = errors.New("invalid comparison")

type temperatureRange struct {
	lower int
	upper int
}

func (allowed *temperatureRange) apply(comparison string, temperature int) (int, error) {
	switch comparison {
	case ">=":
		allowed.lower = max(allowed.lower, temperature)
	case "<=":
		allowed.upper = min(allowed.upper, temperature)
	default:
		return 0, fmt.Errorf("%w: %q", errInvalidComparison, comparison)
	}

	if allowed.lower > allowed.upper {
		return -1, nil
	}

	return allowed.lower, nil
}
