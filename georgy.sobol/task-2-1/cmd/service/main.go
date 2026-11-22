package main

import "fmt"

const (
	maxTemperature = 30
	minTemperature = 15
)

func getComfortableTemperature(sign string, minCurrentT, maxCurrentT *int) int {
	switch sign {
	case "<=":
		var currentTemperature int

		if _, err := fmt.Scan(&currentTemperature); err != nil {
			fmt.Println("Invalid temperature")
		} else if currentTemperature < *maxCurrentT {
			*maxCurrentT = currentTemperature
		}
	case ">=":
		var currentTemperature int

		if _, err := fmt.Scan(&currentTemperature); err != nil {
			fmt.Println("Invalid temperature")
		} else if currentTemperature > *minCurrentT {
			*minCurrentT = currentTemperature
		}
	default:
		fmt.Println("Invalid sign")

		return -1
	}

	if *minCurrentT > *maxCurrentT {
		return -1
	}

	return *minCurrentT
}

func main() {
	var numOfDepartments int
	if _, err := fmt.Scan(&numOfDepartments); err != nil || (numOfDepartments < 1 || numOfDepartments > 1000) {
		fmt.Println("Invalid department number")

		return
	}

	for range numOfDepartments {
		var numOfWorkers int
		if _, err := fmt.Scan(&numOfWorkers); err != nil || (numOfWorkers < 1 || numOfWorkers > 1000) {
			fmt.Println("Invalid workers number")

			return
		}

		minCurrentT := minTemperature
		maxCurrentT := maxTemperature

		for range numOfWorkers {
			var sign string

			if _, err := fmt.Scan(&sign); err != nil {
				fmt.Println("Invalid sign")

				return
			}

			current := getComfortableTemperature(sign, &minCurrentT, &maxCurrentT)
			fmt.Println(current)
		}
	}
}
