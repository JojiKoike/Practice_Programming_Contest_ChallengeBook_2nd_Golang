package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// V is number of vertexes
var V int

// G is adjacency list
var G [][]int

var color []int

func dfs(v int, c int) bool {
	color[v] = c
	for i := 0; i < len(G[v]); i++ {
		// In case adjacent vertex is same color
		if color[G[v][i]] == c {
			return false
		}
		// If adjacent vertex is not yet colored, color it with -c
		if color[G[v][i]] == 0 && !dfs(G[v][i], -c) {
			return false
		}
	}
	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get the number of vertexes
	sc.Scan()
	V, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get connection info and build adjacency list
	G = make([][]int, V)
	for i := 0; i < V; i++ {
		sc.Scan()
		txtconinfo := strings.Split(strings.TrimSpace(sc.Text()), " ")
		s, _ := strconv.Atoi(txtconinfo[0])
		t, _ := strconv.Atoi(txtconinfo[1])
		G[s] = append(G[s], t)
		G[t] = append(G[t], s)
	}

	// Initialize Color List
	color = make([]int, V)

	// Solve
	for i := 0; i < V; i++ {
		if color[i] == 0 {
			if !dfs(i, 1) {
				fmt.Println("No")
				return
			}
		}
	}
	// Show Result
	fmt.Println("Yes")
}
