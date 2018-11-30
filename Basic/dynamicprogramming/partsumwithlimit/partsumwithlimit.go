package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Inputs
var n int
var K int
var a []int
var m []int

// DP Table
var dp [][]bool

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get n
	sc.Scan()
	n, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))
	// Get a, m
	a = []int{}
	m = []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		am := strings.Split(strings.TrimSpace(sc.Text()), " ")
		inta, _ := strconv.Atoi(am[0])
		intm, _ := strconv.Atoi(am[1])
		a = append(a, inta)
		m = append(m, intm)
	}
	fmt.Println(a)
	fmt.Println(m)
	// Get K
	sc.Scan()
	K, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Initialize DP table
	dp = [][]bool{}
	for i := 0; i < n+1; i++ {
		dprow := []bool{}
		for j := 0; j < K+1; j++ {
			dprow = append(dprow, false)
		}
		dp = append(dp, dprow)
	}

	// Solve
	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j <= K; j++ {
			for k := 0; k <= m[i] && k*a[i] <= j; k++ {
				dp[i+1][j] = dp[i+1][j] || dp[i][j-k*a[i]]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= K; j++ {
			fmt.Printf("%v ", dp[i][j])
		}
		fmt.Print("\n")
	}

	if dp[n][K] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
