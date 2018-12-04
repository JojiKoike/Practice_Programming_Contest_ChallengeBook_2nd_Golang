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

	// Get n, m, M
	sc.Scan()
	nmM := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(nmM[0])
	m, _ := strconv.Atoi(nmM[1])
	M, _ := strconv.Atoi(nmM[2])
	fmt.Println(M)

	// Initialize DP Table
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// Solve
	dp[0][0] = 1
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			fmt.Printf("i : %d, j : %d \n", i, j)
			if j-i >= 0 {
				dp[i][j] = (dp[i-1][j] + dp[i][j-i]) % M
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Printf("Result : %d\n", dp[m][n])
}
