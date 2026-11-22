package main

import (
	"fmt"
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

		minTemp := 15
		maxTemp := 30

		for range employeesCount {
			var operator string

			var targetTemp int

			if _, err := fmt.Scan(&operator, &targetTemp); err != nil {
				return
			}

			if operator == ">=" {
				if targetTemp > minTemp {
					minTemp = targetTemp
				}
			} else if operator == "<=" {
				if targetTemp < maxTemp {
					maxTemp = targetTemp
				}
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
