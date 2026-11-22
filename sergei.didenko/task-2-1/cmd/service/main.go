package main

import "fmt"

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for range departments {
		minTemp := 15
		maxTemp := 30

		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			return
		}

		for range employees {
			var operator string

			var tempEmp int
			if _, err := fmt.Scan(&operator, &tempEmp); err != nil {
				return
			}

			switch operator {
			case ">=":
				if tempEmp > minTemp {
					minTemp = tempEmp
				}
			case "<=":
				if tempEmp < maxTemp {
					maxTemp = tempEmp
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
