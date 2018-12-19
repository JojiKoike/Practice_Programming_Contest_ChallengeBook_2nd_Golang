package main

import (
	"bufio"
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
}
