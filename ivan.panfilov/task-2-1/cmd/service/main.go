package main

import "fmt"

func processDepartment(workersCount int) {
	minTemp := 15
	maxTemp := 30

	for range workersCount {
		var operator string

		var temp int

		if _, err := fmt.Scan(&operator, &temp); err != nil {
			fmt.Println("Invalid temperature setting")

			return
		}

		switch operator {
		case ">=":
			if temp > minTemp {
				minTemp = temp
			}
		case "<=":
			if temp < maxTemp {
				maxTemp = temp
			}
		default:
			fmt.Println("Invalid operator")

			continue
		}

		if minTemp > maxTemp {
			fmt.Println(-1)
		} else {
			fmt.Println(minTemp)
		}
	}
}

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid department number")

		return
	}

	for range departmentsCount {
		var workersCount int

		if _, err := fmt.Scan(&workersCount); err != nil {
			fmt.Println("Invalid worker number")

			return
		}

		processDepartment(workersCount)
	}
}
