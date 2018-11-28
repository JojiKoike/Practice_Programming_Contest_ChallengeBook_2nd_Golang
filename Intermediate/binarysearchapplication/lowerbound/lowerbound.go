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

	// Get n, k
	sc.Scan()
	nk := strings.Split(strings.TrimSpace(sc.Text()), " ")
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])

	// Get num array
	a := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		inputa, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
		a = append(a, inputa)
	}

	// Solve
	lb := -1 // Lower Bound
	ub := n  // Upper Bound

	for ub-lb > 1 {
		mid := (lb + ub) / 2
		if a[mid] >= k {
			ub = mid
		} else {
			lb = mid
		}
	}

	fmt.Printf("Result: %d\n", ub)
}
