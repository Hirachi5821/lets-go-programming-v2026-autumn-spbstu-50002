package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const maximumCount = 1000

var errInvalidCount = errors.New("count must be between 1 and 1000")

func readCount(input *bufio.Reader) (int, error) {
	var count int

	if _, err := fmt.Fscan(input, &count); err != nil {
		return 0, fmt.Errorf("read count: %w", err)
	}

	if count < 1 || count > maximumCount {
		return 0, errInvalidCount
	}

	return count, nil
}

func processDepartment(input *bufio.Reader, output *bufio.Writer) error {
	employees, err := readCount(input)
	if err != nil {
		return err
	}

	allowed := temperatureRange{lower: minimumTemperature, upper: maximumTemperature}

	for range employees {
		var (
			comparison  string
			temperature int
		)

		if _, err := fmt.Fscan(input, &comparison, &temperature); err != nil {
			return fmt.Errorf("read temperature constraint: %w", err)
		}

		answer, err := allowed.apply(comparison, temperature)
		if err != nil {
			return err
		}

		if _, err := fmt.Fprintln(output, answer); err != nil {
			return fmt.Errorf("write temperature: %w", err)
		}
	}

	return nil
}

func run(input *bufio.Reader, output *bufio.Writer) error {
	departments, err := readCount(input)
	if err != nil {
		return err
	}

	for range departments {
		if err := processDepartment(input, output); err != nil {
			return err
		}
	}

	if err := output.Flush(); err != nil {
		return fmt.Errorf("flush output: %w", err)
	}

	return nil
}

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)

	if err := run(input, output); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
