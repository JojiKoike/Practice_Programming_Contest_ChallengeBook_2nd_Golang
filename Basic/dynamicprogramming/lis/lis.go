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

	// Get n
	sc.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get a
	a := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		_a, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
		a[i] = _a
	}

	// Initialize DP Table
	dp := make([]int, n)

	// Solve
	var res = 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[i] + 1
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	fmt.Printf("Result : %d\n", res)
}
