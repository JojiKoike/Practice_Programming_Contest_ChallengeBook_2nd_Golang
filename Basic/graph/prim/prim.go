package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cost [][]int
var mincost []int
var used []bool

// INF is infinite value
const INF = 1000000000

// V is number of vertexes
var V int

// E is number of edges
var E int

func prim() int {
	mincost[0] = 0
	var res = 0

	for {
		var v = -1

		for u := 0; u < V; u++ {
			if !used[u] && (v == -1 || mincost[u] < mincost[v]) {
				v = u
			}
		}

		if v == -1 {
			break
		}
		used[v] = true
		res += mincost[v]

		for u := 0; u < V; u++ {
			if mincost[u] > cost[v][u] {
				mincost[u] = cost[v][u]
			}
		}
	}
	return res
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get V
	sc.Scan()
	V, _ = strconv.Atoi(sc.Text())

	// Get E
	sc.Scan()
	E, _ = strconv.Atoi(sc.Text())

	// Initialize Cost
	cost = make([][]int, V)
	mincost = make([]int, V)
	used = make([]bool, V)
	for i := 0; i < V; i++ {
		row := make([]int, V)
		for j := 0; j < V; j++ {
			row[j] = INF
			mincost[j] = INF
			used[j] = false
		}
		cost[i] = row
	}

	// Get Cost Data
	for i := 0; i < E; i++ {
		sc.Scan()
		uvc := strings.Split(sc.Text(), " ")
		u, _ := strconv.Atoi(uvc[0])
		v, _ := strconv.Atoi(uvc[1])
		c, _ := strconv.Atoi(uvc[2])
		cost[u][v] = c
		cost[v][u] = c
	}

	// Solve And Show Result
	fmt.Printf("Result : %d\n", prim())
}
