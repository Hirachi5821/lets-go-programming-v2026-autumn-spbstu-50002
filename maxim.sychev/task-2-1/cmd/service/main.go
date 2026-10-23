package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	minDepts = 1
	maxDepts = 1000
	minEmps  = 1
	maxEmps  = 1000
	minTemp  = 15
	maxTemp  = 30
)

var (
	errReadInt      = errors.New("read int")
	errReadN        = errors.New("read n")
	errReadK        = errors.New("read k")
	errNRange       = errors.New("n out of range")
	errKRange       = errors.New("k out of range")
	errReadOperator = errors.New("unexpected end of input")
	errUnknownOp    = errors.New("unknown operator")
	errReadTemp     = errors.New("expected temperature")
	errWriteLine    = errors.New("write line")
)

func readInt(reader *bufio.Reader) (int, error) {
	var value int

	if _, err := fmt.Fscan(reader, &value); err != nil {
		return 0, fmt.Errorf("%w: %w", errReadInt, err)
	}

	return value, nil
}

func readDeptCount(reader *bufio.Reader) (int, error) {
	value, err := readInt(reader)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", errReadN, err)
	}

	if value < minDepts || value > maxDepts {
		return 0, fmt.Errorf("%w: N=%d, range [%d..%d]",
			errNRange, value, minDepts, maxDepts)
	}

	return value, nil
}

func readEmpCount(reader *bufio.Reader, dept int) (int, error) {
	value, err := readInt(reader)
	if err != nil {
		return 0, fmt.Errorf("%w: department %d: %w", errReadK, dept, err)
	}

	if value < minEmps || value > maxEmps {
		return 0, fmt.Errorf("%w: K=%d, department %d, range [%d..%d]",
			errKRange, value, dept, minEmps, maxEmps)
	}

	return value, nil
}

func readConstraint(
	reader *bufio.Reader,
	dept, employee, total int,
) (string, int, error) {
	var operator string

	if _, err := fmt.Fscan(reader, &operator); err != nil {
		return "", 0, fmt.Errorf("%w: department %d, employee %d/%d: %w",
			errReadOperator, dept, employee, total, err)
	}

	if operator != ">=" && operator != "<=" {
		return "", 0, fmt.Errorf("%w: %q, department %d, employee %d",
			errUnknownOp, operator, dept, employee)
	}

	value, err := readInt(reader)
	if err != nil {
		return "", 0, fmt.Errorf("%w after %q: department %d, employee %d/%d: %w",
			errReadTemp, operator, dept, employee, total, err)
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

func printLine(writer *bufio.Writer, s string) error {
	if _, err := fmt.Fprintln(writer, s); err != nil {
		return fmt.Errorf("%w: %w", errWriteLine, err)
	}

	return nil
}

func printResult(writer *bufio.Writer, low, high int) error {
	if low <= high {
		return printLine(writer, strconv.Itoa(low))
	}

	return printLine(writer, "-1")
}

func processDepartment(reader *bufio.Reader, writer *bufio.Writer, dept int) error {
	employees, err := readEmpCount(reader, dept)
	if err != nil {
		return err
	}

	low, high := minTemp, maxTemp

	for emp := range employees {
		operator, value, consErr := readConstraint(reader, dept, emp+1, employees)
		if consErr != nil {
			return consErr
		}

		low, high = applyConstraint(low, high, operator, value)

		if err = printResult(writer, low, high); err != nil {
			return err
		}
	}

	return nil
}

func logError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, "error:", err)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer func() {
		_ = writer.Flush()
	}()

	departments, err := readDeptCount(reader)
	if err != nil {
		logError(err)

		return
	}

	for dept := range departments {
		if err = processDepartment(reader, writer, dept+1); err != nil {
			logError(err)

			return
		}
	}
}
