package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// N
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	// Matrix
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		row := make([]int, N)
		rowinputs := strings.Split(sc.Text(), "")
		for j, rowinput := range rowinputs {
			a, _ := strconv.Atoi(rowinput)
			row[j] = a
		}
		A[i] = row
	}

	// "1" last position in row
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = -1
		for j := 0; j < N; j++ {
			if A[i][j] == 1 {
				a[i] = j
			}
		}
	}

	// Solve
	res := 0
	for i := 0; i < N; i++ {
		pos := -1
		for j := i; j < N; j++ {
			if a[j] <= i {
				pos = j
				break
			}
		}

		for j := pos; j > i; j-- {
			a[j], a[j-1] = a[j-1], a[j]
			res++
		}
	}

	// Show Results
	fmt.Printf("Result: %d\n", res)
}
