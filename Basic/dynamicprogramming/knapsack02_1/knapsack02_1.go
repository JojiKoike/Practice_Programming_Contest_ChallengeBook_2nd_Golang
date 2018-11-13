package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type item struct {
	w, v int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Number of items
	sc.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get Items Data
	items := []item{}
	for i := 0; i < n; i++ {
		sc.Scan()
		txtinput := strings.Split(strings.TrimSpace(sc.Text()), " ")
		weight, _ := strconv.Atoi(txtinput[0])
		value, _ := strconv.Atoi(txtinput[1])
		items = append(items, item{w: weight, v: value})
	}
	// Available Weight
	sc.Scan()
	W, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Initialize DP Table
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, W+1)
	}

	// Solve
	for i := 0; i < n; i++ {
		for j := 0; j <= W; j++ {
			if j < items[i].w {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] =
					int(math.Max(float64(dp[i][j]),
						float64(dp[i+1][j-items[i].w]+items[i].v)))
			}
			fmt.Print(dp[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Printf("Result %d\n", dp[n][W])
}
