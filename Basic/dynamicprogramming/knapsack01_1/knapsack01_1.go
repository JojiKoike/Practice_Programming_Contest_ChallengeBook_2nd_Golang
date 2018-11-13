package main

import (
	"bufio"
	"fmt"
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

// DP Table
var dp = [][]int{}

func rec(i, j int) int {
	if dp[i][j] >= 0 {
		return dp[i][j]
	}
	var res int
	if i == N {
		res = 0
	} else if j < items[i].w {
		res = rec(i+1, j)
	} else {
		rec1 := rec(i+1, j)
		rec2 := rec(i+1, j-items[i].w) + items[i].v
		if rec1 < rec2 {
			res = rec2
		} else {
			res = rec1
		}
	}
	dp[i][j] = res
	return res
}

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
	for i := 0; i < N+1; i++ {
		row := []int{}
		for j := 0; j < W+1; j++ {
			row = append(row, -1)
		}
		dp = append(dp, row)
	}

	// Calc And Show Result
	fmt.Printf("Result : %d\n", rec(0, W))
}
