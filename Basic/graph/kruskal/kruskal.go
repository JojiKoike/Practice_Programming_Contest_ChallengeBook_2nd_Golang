package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Edge
type edge struct {
	u, v, cost int
}
type edges []edge

func (e edges) Len() int           { return len(e) }
func (e edges) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e edges) Less(i, j int) bool { return e[i].cost < e[j].cost }

var es edges

// Union Find Tree
var par []int
var rank []int

func initUnionFind(v int) {
	par = make([]int, v)
	rank = make([]int, v)
	for i := 0; i < v; i++ {
		par[i] = i
		rank[i] = 0
	}
}
func find(x int) int {
	if par[x] == x {
		return x
	}
	par[x] = find(par[x])
	return par[x]
}
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
func same(x int, y int) bool {
	return find(x) == find(y)
}

// V is number of vertexes
var V int

// E is number of edges
var E int

func kruskal() int {
	sort.Sort(edges(es))
	initUnionFind(V)
	var res = 0
	for i := 0; i < 2*E; i++ {
		e := es[i]
		if !same(e.u, e.v) {
			unite(e.u, e.v)
			res += e.cost
		}
	}
	return res
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// V
	sc.Scan()
	V, _ = strconv.Atoi(sc.Text())

	// E
	sc.Scan()
	E, _ = strconv.Atoi(sc.Text())

	// Edge
	es = make([]edge, 2*E)
	for i := 0; i < E; i++ {
		sc.Scan()
		uvc := strings.Split(sc.Text(), " ")
		_u, _ := strconv.Atoi(uvc[0])
		_v, _ := strconv.Atoi(uvc[1])
		_cost, _ := strconv.Atoi(uvc[2])
		es[i] = edge{u: _u, v: _v, cost: _cost}
		es[i+E-1] = edge{u: _v, v: _u, cost: _cost}
	}

	fmt.Printf("Result: %d\n", kruskal())

}
