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

// N is Number of items
var N int

// W is Available Total Weight
var W int

// Item List
var items = []item{}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get Number of Items
	sc.Scan()
	N, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get Items Data
	for i := 0; i < N; i++ {
		sc.Scan()
		txtitem := strings.Split(strings.TrimSpace(sc.Text()), " ")
		weight, _ := strconv.Atoi(txtitem[0])
		value, _ := strconv.Atoi(txtitem[1])
		items = append(items, item{w: weight, v: value})
	}

	// Get Available Total Weight
	sc.Scan()
	W, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Initialize DP Table
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := N - 1; i >= 0; i-- {
		for j := 0; j <= W; j++ {
			if j < items[i].w {
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] =
					int(math.Max(float64(dp[i+1][j]),
						float64(dp[i+1][j-items[i].w]+items[i].v)))
			}
		}
	}
	fmt.Printf("Result : %d\n", dp[0][W])

}
