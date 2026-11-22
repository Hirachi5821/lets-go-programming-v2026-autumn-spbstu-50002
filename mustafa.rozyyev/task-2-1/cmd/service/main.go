package main

import (
	"errors"
	"fmt"
)

const (
	minCount = 1
	maxCount = 1000

	minTemp = 15
	maxTemp = 30
)

var (
	errInvalidNumber   = errors.New("invalid number")
	errInvalidOperator = errors.New("invalid operator")
	errInvalidTemp     = errors.New("invalid temperature")
)

type temperatureConstraint struct {
	operator    string
	temperature int
}

func readAndCheckCount() (int, error) {
	var count int
	if _, err := fmt.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to scan count: %w", err)
	}

	if count < minCount || count > maxCount {
		return 0, fmt.Errorf("%w: must be between %d and %d", errInvalidNumber, minCount, maxCount)
	}

	return count, nil
}

func readInput() ([][]temperatureConstraint, error) {
	departmentCount, err := readAndCheckCount()
	if err != nil {
		return nil, err
	}

	departments := make([][]temperatureConstraint, departmentCount)

	for departmentIndex := range departmentCount {
		employeeCount, err := readAndCheckCount()
		if err != nil {
			return nil, err
		}

		departments[departmentIndex] = make([]temperatureConstraint, employeeCount)

		for employeeIndex := range employeeCount {
			var constraint temperatureConstraint
			if _, err := fmt.Scan(&constraint.operator, &constraint.temperature); err != nil {
				return nil, fmt.Errorf("failed to scan temperature request: %w", err)
			}

			if constraint.operator != ">=" && constraint.operator != "<=" {
				return nil, fmt.Errorf("%w: %q", errInvalidOperator, constraint.operator)
			}

			if constraint.temperature < minTemp || constraint.temperature > maxTemp {
				return nil, fmt.Errorf("%w: must be between %d and %d", errInvalidTemp, minTemp, maxTemp)
			}

			departments[departmentIndex][employeeIndex] = constraint
		}
	}

	return departments, nil
}

func processDepartment(constraints []temperatureConstraint) []int {
	results := make([]int, 0, len(constraints))
	lower, upper := minTemp, maxTemp

	for _, constraint := range constraints {
		switch constraint.operator {
		case ">=":
			if constraint.temperature > lower {
				lower = constraint.temperature
			}
		case "<=":
			if constraint.temperature < upper {
				upper = constraint.temperature
			}
		}

		if lower <= upper {
			results = append(results, lower)
		} else {
			results = append(results, -1)
		}
	}

	return results
}

func main() {
	departments, err := readInput()
	if err != nil {
		fmt.Println(err)

		return
	}

	for _, department := range departments {
		for _, result := range processDepartment(department) {
			fmt.Println(result)
		}
	}
}
