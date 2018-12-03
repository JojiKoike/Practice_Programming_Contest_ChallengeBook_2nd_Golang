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

	// Get n (Number of items) and K (Total Summation)
	sc.Scan()
	nK := strings.Split(strings.TrimSpace(sc.Text()), " ")
	n, _ := strconv.Atoi(nK[0])
	K, _ := strconv.Atoi(nK[1])

	// Get items data (value, number of )
	a := make([]int, n)
	m := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		am := strings.Split(strings.TrimSpace(sc.Text()), " ")
		_a, _ := strconv.Atoi(am[0])
		_m, _ := strconv.Atoi(am[1])
		a[i] = _a
		m[i] = _m
	}

	// Initialize DP Table
	dp := make([]int, K+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0

	// Solve
	for i := 0; i < n; i++ {
		for j := 0; j <= K; j++ {
			if dp[j] >= 0 {
				dp[j] = m[i]
			} else if j < a[i] || dp[j-a[i]] <= 0 {
				dp[j] = -1
			} else {
				dp[j] = dp[j-a[i]] - 1
			}
		}
	}

	// Show Result
	if dp[K] >= 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
