package main

import (
	"fmt"
)

func main() {
	var departmentsCount, employeesCount, temperature int

	var symbol string

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		if _, err := fmt.Scan(&employeesCount); err != nil {
			return
		}

		minT := 15
		maxT := 30

		for range employeesCount {
			if _, err := fmt.Scan(&symbol, &temperature); err != nil {
				return
			}

			switch symbol {
			case "<=":
				if temperature < maxT {
					maxT = temperature
				}
			case ">=":
				if temperature > minT {
					minT = temperature
				}
			default:
				fmt.Print("Wrong symbol\n")
			}

			if maxT >= minT {
				fmt.Println(minT)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
