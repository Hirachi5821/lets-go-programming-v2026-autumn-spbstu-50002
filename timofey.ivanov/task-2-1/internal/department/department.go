package department

import (
	"errors"
	"fmt"
)

var ErrInvalidOperator = errors.New("invalid operator")

type Department struct {
	minTemp int
	maxTemp int
}

func NewDepartment(minTemp, maxTemp int) *Department {
	return &Department{
		minTemp: minTemp,
		maxTemp: maxTemp,
	}
}

func (d *Department) ApplyConstraint(operator string, temp int) error {
	switch operator {
	case ">=":
		if temp > d.minTemp {
			d.minTemp = temp
		}
	case "<=":
		if temp < d.maxTemp {
			d.maxTemp = temp
		}
	default:
		return fmt.Errorf("%w: %q", ErrInvalidOperator, operator)
	}

	return nil
}

func (d *Department) OptimalTemp() int {
	if d.minTemp > d.maxTemp {
		return -1
	}

	return d.minTemp
}
