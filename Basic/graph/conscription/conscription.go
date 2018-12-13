package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type edge struct {
	u, v, cost int
}

// Define Sortable Ascending by cost List
type edges []edge

func (e edges) Len() int           { return len(e) }
func (e edges) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e edges) Less(i, j int) bool { return e[i].cost < e[j].cost }

var es edges

//---- Define Union Find Tree Methods ---- //
var par []int
var rank []int

// Initialize Union Find Tree
func initUnionFind(v int) {
	par = make([]int, v)
	rank = make([]int, v)
	for i := 0; i < v; i++ {
		par[i] = i
		rank[i] = 0
	}
}

// Find Root Parent
func find(x int) int {
	if par[x] == x {
		return x
	}
	par[x] = find(par[x])
	return par[x]
}

// Unite Tree
func unite(x int, y int) {
	x = find(x)
	y = find(y)
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

// Check Same Group
func same(x int, y int) bool {
	return find(x) == find(y)
}

// V is number of vertexes
var V int

// E is number of edges
var E int

// MST Solver (Kruskal Method)
func kruskal() int {
	sort.Sort(edges(es))
	initUnionFind(V)
	var res = 0
	for i := 0; i < E; i++ {
		e := es[i]
		if !same(e.u, e.v) {
			unite(e.u, e.v)
			res += e.cost
		}
	}
	return res
}

// N is number of man
var N int

// M is number of women
var M int

// R is number of relations
var R int

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get N (man), M (women), R (relations)
	sc.Scan()
	nmr := strings.Split(sc.Text(), " ")
	N, _ = strconv.Atoi(nmr[0])
	M, _ = strconv.Atoi(nmr[1])
	R, _ = strconv.Atoi(nmr[2])
	E = R
	V = N + R

	// Get Relations Data
	es = make([]edge, E)
	for i := 0; i < R; i++ {
		sc.Scan()
		xyd := strings.Split(sc.Text(), " ")
		x, _ := strconv.Atoi(xyd[0])
		y, _ := strconv.Atoi(xyd[1])
		d, _ := strconv.Atoi(xyd[2])
		es[i] = edge{u: x, v: M + y, cost: -d}
	}

	// Calc And Show Result
	fmt.Printf("Result %d\n", 10000*(N+M)+kruskal())
}
