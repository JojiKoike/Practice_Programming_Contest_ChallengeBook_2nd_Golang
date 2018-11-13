package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
	sc.Scan()
	m, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
	sc.Scan()
	s := strings.Split(strings.TrimSpace(sc.Text()), "")
	sc.Scan()
	t := strings.Split(strings.TrimSpace(sc.Text()), "")

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = int(math.Max(float64(dp[i][j+1]), float64(dp[i+1][j])))
			}
		}
	}
	fmt.Printf("Result %d\n", dp[n][m])
}
