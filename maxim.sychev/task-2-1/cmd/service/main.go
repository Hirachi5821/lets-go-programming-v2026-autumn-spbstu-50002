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

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		fmt.Fprintln(out, "error: failed to read N:", err)
		return
	}
	if n < minDepts || n > maxDepts {
		fmt.Fprintf(out, "error: N=%d out of range [%d..%d]\n", n, minDepts, maxDepts)
		return
	}

	for dept := 0; dept < n; dept++ {
		var k int
		if _, err := fmt.Fscan(in, &k); err != nil {
			fmt.Fprintf(out, "error: failed to read K for department %d: %v\n", dept+1, err)
			return
		}
		if k < minEmps || k > maxEmps {
			fmt.Fprintf(out, "error: K=%d for department %d out of range [%d..%d]\n",
				k, dept+1, minEmps, maxEmps)
			return
		}

		low, high := minTemp, maxTemp

		for emp := 0; emp < k; emp++ {
			var op string
			var t int
			if _, err := fmt.Fscan(in, &op, &t); err != nil {
				fmt.Fprintf(out, "error: unexpected end of input (department %d, employee %d/%d): %v\n",
					dept+1, emp+1, k, err)
				return
			}

			switch op {
			case ">=":
				if t > low {
					low = t
				}
			case "<=":
				if t < high {
					high = t
				}
			default:
				fmt.Fprintf(out, "error: unknown operator %q (department %d, employee %d)\n",
					op, dept+1, emp+1)
				return
			}

			if low <= high {
				fmt.Fprintln(out, low)
			} else {
				fmt.Fprintln(out, -1)
			}
		}
	}
}
