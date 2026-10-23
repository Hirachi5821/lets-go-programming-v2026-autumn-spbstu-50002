package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	minDepts = 1
	maxDepts = 1000
	minEmps  = 1
	maxEmps  = 1000
	minTemp  = 15
	maxTemp  = 30
)

func readInt(r *bufio.Reader) (int, error) {
	var value int

	if _, err := fmt.Fscan(r, &value); err != nil {
		return 0, fmt.Errorf("read int: %w", err)
	}

	return value, nil
}

func readDeptCount(r *bufio.Reader) (int, error) {
	value, err := readInt(r)
	if err != nil {
		return 0, fmt.Errorf("read N: %w", err)
	}

	if value < minDepts || value > maxDepts {
		return 0, fmt.Errorf("N=%d out of range [%d..%d]", value, minDepts, maxDepts)
	}

	return value, nil
}

func readEmpCount(r *bufio.Reader, dept int) (int, error) {
	value, err := readInt(r)
	if err != nil {
		return 0, fmt.Errorf("read K for department %d: %w", dept, err)
	}

	if value < minEmps || value > maxEmps {
		return 0, fmt.Errorf("K=%d for department %d out of range [%d..%d]",
			value, dept, minEmps, maxEmps)
	}

	return value, nil
}

func readConstraint(r *bufio.Reader, dept, emp, total int) (string, int, error) {
	var operator string

	if _, err := fmt.Fscan(r, &operator); err != nil {
		return "", 0, fmt.Errorf(
			"unexpected end of input (department %d, employee %d/%d): %w",
			dept, emp, total, err,
		)
	}

	if operator != ">=" && operator != "<=" {
		return "", 0, fmt.Errorf("unknown operator %q (department %d, employee %d)",
			operator, dept, emp)
	}

	value, err := readInt(r)
	if err != nil {
		return "", 0, fmt.Errorf("expected temperature after %q (department %d, employee %d/%d): %w",
			operator, dept, emp, total, err)
	}

	return operator, value, nil
}

func applyConstraint(low, high int, operator string, value int) (int, int) {
	switch operator {
	case ">=":
		if value > low {
			low = value
		}
	case "<=":
		if value < high {
			high = value
		}
	}

	return low, high
}

func printLine(w *bufio.Writer, s string) error {
	if _, err := fmt.Fprintln(w, s); err != nil {
		return fmt.Errorf("write line: %w", err)
	}

	return nil
}

func processDepartment(r *bufio.Reader, w *bufio.Writer, dept, total int) error {
	employees, err := readEmpCount(r, dept)
	if err != nil {
		return err
	}

	low, high := minTemp, maxTemp

	for emp := range employees {
		operator, value, consErr := readConstraint(r, dept, emp+1, employees)
		if consErr != nil {
			return consErr
		}

		low, high = applyConstraint(low, high, operator, value)

		if low <= high {
			if err = printLine(w, fmt.Sprintf("%d", low)); err != nil {
				return err
			}
		} else {
			if err = printLine(w, "-1"); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer func() {
		_ = writer.Flush()
	}()

	departments, err := readDeptCount(reader)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)

		return
	}

	for dept := range departments {
		if err = processDepartment(reader, writer, dept+1, departments); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)

			return
		}
	}
}
