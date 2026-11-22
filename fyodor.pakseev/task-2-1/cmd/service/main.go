package main

import (
	"fmt"
)

const (
	defMinTemp = 15
	defMaxTemp = 30
)

func addTemp(operation string, value int, minTemp, maxTemp int) (int, int) {
	switch operation {
	case "<=":
		if value < maxTemp {
			maxTemp = value
		}
	case ">=":
		if value > minTemp {
			minTemp = value
		}
	}

	return minTemp, maxTemp
}

func main() {
	var departs int

	if _, err := fmt.Scan(&departs); err != nil {
		return
	}

	for range departs {
		var numStaff int

		if _, err := fmt.Scan(&numStaff); err != nil {
			return
		}

		minTemp, maxTemp := defMinTemp, defMaxTemp

		for range numStaff {
			var (
				operation string
				value     int
			)

			if _, err := fmt.Scan(&operation, &value); err != nil {
				return
			}

			minTemp, maxTemp = addTemp(operation, value, minTemp, maxTemp)

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
