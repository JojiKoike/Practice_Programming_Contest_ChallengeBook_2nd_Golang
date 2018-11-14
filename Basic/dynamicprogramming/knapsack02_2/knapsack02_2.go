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

	// Initialize DP Array
	dp := make([]int, W+1)

	// Solve
	for i := 0; i < n; i++ {
		for j := items[i].w; j <= W; j++ {
			dp[j] = int(math.Max(
				float64(dp[j]),
				float64(dp[j-items[i].w]+items[i].v)))
		}
	}

	// Show Result
	fmt.Printf("Result : %d\n", dp[W])
}
