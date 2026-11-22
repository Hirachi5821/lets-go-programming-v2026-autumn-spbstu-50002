package main

import "fmt"

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Error: Enter a whole number for the departments")

		return
	}

	if departmentCount < 1 || departmentCount > 1000 {
		fmt.Println("Error: the number of departments must be between 1 and 1000")

		return
	}

	for i := 1; i <= departmentCount; i++ {
		var employeeCount int

		minTemp := 15
		maxTemp := 30

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Error: Enter a whole number for the employees")

			return
		}

		if employeeCount < 1 || employeeCount > 1000 {
			fmt.Println("Error: the number of employees must be between 1 and 1000")

			return
		}

		for j := 1; j <= employeeCount; j++ {
			var (
				operation string
				temp      int
			)

			if _, err := fmt.Scan(&operation, &temp); err != nil {
				fmt.Println("Error: Enter an operator (>= or <=) and a whole number for the temp")

				return
			}

			if !updateTemperature(operation, temp, &minTemp, &maxTemp) {
				fmt.Println("Error: operator must be >= or <=")

				return
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

func updateTemperature(operation string, temp int, minTemp, maxTemp *int) bool {
	switch operation {
	case ">=":
		if temp > *minTemp {
			*minTemp = temp
		}
	case "<=":
		if temp < *maxTemp {
			*maxTemp = temp
		}
	default:
		return false
	}

	return true
}
