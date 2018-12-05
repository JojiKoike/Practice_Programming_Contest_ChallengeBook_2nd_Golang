package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Parent
var par []int

// Tree Depth
var rank []int

// Initialize
func initialize(n int) {
	par = make([]int, n)
	rank = make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
		rank[i] = 0
	}
}

// Calc Tree Root
func find(x int) int {
	if par[x] == x {
		return x
	}
	par[x] = find(par[x])
	return par[x]
}

// Marge Groups x and y belongs to
func unite(x int, y int) {
	x = find(x)
	y = find(y)
	// Same root = x and y belongs to same group
	if x == y {
		return
	}
	if rank[x] < rank[y] {
		par[x] = y
	} else {
		par[y] = x
		if rank[x] == rank[y] {
			rank[x]++
		}
	}
}

// Check x and y belongs to same group
func same(x int, y int) bool {
	return find(x) == find(y)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get Number of Elements
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// Initialize Union Find Tree
	initialize(n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d root : %d\n", i, find(i))
	}

	// Merge
	unite(1, 2)
	for i := 0; i < n; i++ {
		fmt.Printf("%d root : %d\n", i, find(i))
	}

	// Judge Same Group
	fmt.Printf("%t\n", same(1, 2))
	fmt.Printf("%t\n", same(0, 1))
}
