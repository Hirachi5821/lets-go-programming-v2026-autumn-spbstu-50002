package main

import (
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			return
		}

		lowerBound, upperBound := minTemperature, maxTemperature

		for range employeesCount {
			var (
				operation   string
				temperature int
			)

			if _, err := fmt.Scan(&operation, &temperature); err != nil {
				return
			}

			switch operation {
			case ">=":
				if temperature > lowerBound {
					lowerBound = temperature
				}
			case "<=":
				if temperature < upperBound {
					upperBound = temperature
				}
			}

			if lowerBound <= upperBound {
				fmt.Println(lowerBound)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
