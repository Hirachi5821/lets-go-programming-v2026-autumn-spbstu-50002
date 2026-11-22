package main

import (
	"fmt"
)

const (
	initialMinTemperature = 15
	initialMaxTemperature = 30
)

type Department struct {
	minTemperature int
	maxTemperature int
}

func NewDepartment() *Department {
	return &Department{
		minTemperature: initialMinTemperature,
		maxTemperature: initialMaxTemperature,
	}
}

func (department *Department) ApplyPreference(operator string, value int) {
	switch operator {
	case ">=":
		if value > department.minTemperature {
			department.minTemperature = value
		}
	case "<=":
		if value < department.maxTemperature {
			department.maxTemperature = value
		}
	}
}

func (department *Department) GetOptimalTemperature() int {
	if department.minTemperature > department.maxTemperature {
		return -1
	}

	return department.minTemperature
}

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for departmentIndex := 0; departmentIndex < departmentsCount; departmentIndex++ {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			return
		}

		department := NewDepartment()

		for employeeIndex := 0; employeeIndex < employeesCount; employeeIndex++ {
			var operator string

			var value int

			if _, err := fmt.Scan(&operator, &value); err != nil {
				return
			}

			department.ApplyPreference(operator, value)

			fmt.Println(department.GetOptimalTemperature())
		}
	}
}
