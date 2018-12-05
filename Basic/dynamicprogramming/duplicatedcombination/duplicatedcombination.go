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

	// Step 1: Get Inputs
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	ainputs := strings.Split(sc.Text(), " ")
	a := make([]int, n)
	for i, ainput := range ainputs {
		a[i], _ = strconv.Atoi(ainput)
	}
	sc.Scan()
	M, _ := strconv.Atoi(sc.Text())

	// Step 2: Initialize DP Table
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		row := make([]int, m+1)
		for j := 0; j < m+1; j++ {
			row[j] = 0
		}
		dp[i] = row
	}

	// Step 3: Solve
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			if j-1-a[i] >= 0 {
				dp[i+1][j] = (dp[i+1][j-1] + dp[i][j] - dp[i][j-1-a[i]]) % M
			} else {
				dp[i+1][j] = (dp[i+1][j-1] + dp[i][j]) % M
			}
		}
	}

	// Step 4: Show Result
	fmt.Printf("Result %d\n", dp[n][m])
}
