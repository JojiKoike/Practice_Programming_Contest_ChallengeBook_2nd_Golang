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

// INF is infinite value for DP table initial value
const INF = 100000000

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get Number of Items
	sc.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get Items Data and calc max value
	items := []item{}
	var maxvalue int
	for i := 0; i < n; i++ {
		sc.Scan()
		txtinput := strings.Split(strings.TrimSpace(sc.Text()), " ")
		weight, _ := strconv.Atoi(txtinput[0])
		value, _ := strconv.Atoi(txtinput[1])
		if value > maxvalue {
			maxvalue = value
		}
		items = append(items, item{w: weight, v: value})
	}

	// Available Weight
	sc.Scan()
	W, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Initialize DP Table
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n*maxvalue+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < n*maxvalue+1; j++ {
			dp[i][j] = INF
		}
	}

	// Create DP Table
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= n*maxvalue; j++ {
			if j < items[i].v {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = int(math.Min(
					float64(dp[i][j]),
					float64(dp[i][j-items[i].v]+items[i].w)))
			}
		}
	}

	// Calc Result
	var res = 0
	for i := 0; i <= n*maxvalue; i++ {
		if dp[n][i] <= W {
			res = i
		}
	}

	fmt.Printf("Result %d\n", res)

}
