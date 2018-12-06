package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Parent
var par []int

// Tree Depth
var rank []int

// Initialize Union Find Tree
func initialize(n int) {
	par = make([]int, n)
	rank = make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
		rank[i] = 0
	}
}

// Find Tree Root
func find(x int) int {
	if par[x] == x {
		return x
	}
	par[x] = find(par[x])
	return par[x]
}

// Unite Groups
func unite(x int, y int) {
	x = find(x)
	y = find(y)
	// If same root, needn't to unite
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

func same(x int, y int) bool {
	return find(x) == find(y)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get N (num of animals) and K (num of info)
	sc.Scan()
	nk := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(nk[0])
	K, _ := strconv.Atoi(nk[1])

	// Get Information
	T := make([]int, K)
	X := make([]int, K)
	Y := make([]int, K)
	for i := 0; i < K; i++ {
		sc.Scan()
		txy := strings.Split(strings.TrimSpace(sc.Text()), " ")
		T[i], _ = strconv.Atoi(txy[0])
		X[i], _ = strconv.Atoi(txy[1])
		Y[i], _ = strconv.Atoi(txy[2])
	}

	// Initialize Union Find Tree
	initialize(3 * N)

	// Solve
	var res = 0
	for i := 0; i < K; i++ {
		t := T[i]
		x := X[i] - 1
		y := Y[i] - 1

		// Incorrect Number
		if x < 0 || N <= x || y < 0 || N <= y {
			return
		}

		if t == 1 {
			// Type1
			if same(x, y+N) || same(x, y+2*N) {
				res++
			} else {
				unite(x, y)
				unite(x+N, y+N)
				unite(x+2*N, y+2*N)
			}
		} else {
			// Type2
			if same(x, y) || same(x, y+2*N) {
				res++
			} else {
				unite(x, y+N)
				unite(x+N, y+2*N)
				unite(x+2*N, y)
			}
		}
	}

	fmt.Printf("Result %d\n", res)

}
