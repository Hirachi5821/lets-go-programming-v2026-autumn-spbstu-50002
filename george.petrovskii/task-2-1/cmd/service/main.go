package main

import (
	"errors"
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30
	minCount       = 1
	maxCount       = 1000
)

var (
	errInvalidCount = errors.New("count must be in range [1,1000]")
	errInvalidSign  = errors.New("sign must be <= or >=")
)

type Department struct {
	minT int
	maxT int
}

func NewDepartment(minTemp, maxTemp int) Department {
	return Department{minT: minTemp, maxT: maxTemp}
}

func (d *Department) AddTemperature(sign string, temp int) (int, error) {
	switch sign {
	case ">=":
		if temp > d.minT {
			d.minT = temp
		}
	case "<=":
		if temp < d.maxT {
			d.maxT = temp
		}
	default:
		return 0, fmt.Errorf("%w: %q", errInvalidSign, sign)
	}

	if d.minT > d.maxT {
		return -1, nil
	}

	return d.minT, nil
}

func readCount() (int, error) {
	var number int

	if _, err := fmt.Scan(&number); err != nil {
		return 0, fmt.Errorf("failed to scan number: %w", err)
	}

	if number < minCount || number > maxCount {
		return 0, errInvalidCount
	}

	return number, nil
}

func main() {
	departments, err := readCount()
	if err != nil {
		fmt.Println(err)

		return
	}

	for range departments {
		staff, err := readCount()
		if err != nil {
			fmt.Println(err)

			return
		}

		department := NewDepartment(minTemperature, maxTemperature)

		for range staff {
			var (
				sign  string
				value int
			)

			if _, err := fmt.Scan(&sign, &value); err != nil {
				fmt.Println("input error:", err)

				return
			}

			result, err := department.AddTemperature(sign, value)
			if err != nil {
				fmt.Println("operation error:", err)

				return
			}

			fmt.Println(result)
		}
	}
}
