package main

import "fmt"

const (
	minAllowedTemperature = 15
	maxAllowedTemperature = 30
)

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for range departments {
		var employees int

		if _, err := fmt.Scan(&employees); err != nil {
			return
		}

		minTemperature := minAllowedTemperature
		maxTemperature := maxAllowedTemperature

		for range employees {
			var (
				comparison  string
				temperature int
			)

			if _, err := fmt.Scan(&comparison, &temperature); err != nil {
				return
			}

			switch comparison {
			case ">=":
				if temperature > minTemperature {
					minTemperature = temperature
				}
			case "<=":
				if temperature < maxTemperature {
					maxTemperature = temperature
				}
			}

			if minTemperature > maxTemperature {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemperature)
			}
		}
	}
}
