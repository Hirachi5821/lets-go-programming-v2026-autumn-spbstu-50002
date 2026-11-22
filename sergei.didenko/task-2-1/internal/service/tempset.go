package service

import (
	"fmt"
	"os"
)

const (
	minInitialTemp = 15
	maxInitialTemp = 30
)

func Run() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения количества отделов:", err)

		return
	}

	for range departments {
		if !processDepartment() {
			return
		}
	}
}

func processDepartment() bool {
	minTemp := minInitialTemp
	maxTemp := maxInitialTemp

	var employees int

	if _, err := fmt.Scan(&employees); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения количества сотрудников:", err)

		return false
	}

	for range employees {
		var isValid bool

		minTemp, maxTemp, isValid = processEmployee(minTemp, maxTemp)
		if !isValid {
			return false
		}

		printTemperature(minTemp, maxTemp)
	}

	return true
}

func processEmployee(minTemp, maxTemp int) (int, int, bool) {
	var operator string

	var tempEmp int

	if _, err := fmt.Scan(&operator, &tempEmp); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения ограничения температуры:", err)

		return minTemp, maxTemp, false
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
	default:
		fmt.Fprintln(os.Stderr, "Ошибка: неверный оператор", operator)

		return minTemp, maxTemp, false
	}

	return minTemp, maxTemp, true
}

func printTemperature(minTemp, maxTemp int) {
	if minTemp <= maxTemp {
		fmt.Println(minTemp)

		return
	}

	fmt.Println(-1)
}
