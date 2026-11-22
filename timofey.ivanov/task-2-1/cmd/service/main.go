package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Kabisdox/task-2-1/internal/department"
)

const (
	minCount = 1
	maxCount = 1000

	departMinTemp = 15
	departMaxTemp = 30
)

var (
	errInvalidNum  = errors.New("invalid number")
	errInvalidTemp = errors.New("invalid temperature")
)

func numChecks() (int, error) {
	var number int
	if _, err := fmt.Scan(&number); err != nil {
		return 0, fmt.Errorf("failed to scan number: %w", err)
	}

	if number < minCount || number > maxCount {
		return 0, fmt.Errorf("%w: must be between %d and %d", errInvalidNum, minCount, maxCount)
	}

	return number, nil
}

func readConstraint() (string, int, error) {
	var operator string
	if _, err := fmt.Scan(&operator); err != nil {
		return "", 0, fmt.Errorf("failed to scan operator: %w", err)
	}

	var temp int
	if _, err := fmt.Scan(&temp); err != nil {
		return "", 0, fmt.Errorf("failed to scan temperature: %w", err)
	}

	if temp < departMinTemp || temp > departMaxTemp {
		return "", 0, fmt.Errorf("%w: must be between %d and %d", errInvalidTemp, departMinTemp, departMaxTemp)
	}

	return operator, temp, nil
}

func processDepart(dept *department.Department, employeesCount int) error {
	for range employeesCount {
		operator, temp, err := readConstraint()
		if err != nil {
			return err
		}

		if err := dept.ApplyConstraint(operator, temp); err != nil {
			return fmt.Errorf("failed to apply constraint: %w", err)
		}

		fmt.Println(dept.OptimalTemp())
	}

	return nil
}

func main() {
	numberOfDepartments, err := numChecks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for range numberOfDepartments {
		numberOfEmployees, err := numChecks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		dept := department.NewDepartment(departMinTemp, departMaxTemp)

		if err := processDepart(dept, numberOfEmployees); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}
