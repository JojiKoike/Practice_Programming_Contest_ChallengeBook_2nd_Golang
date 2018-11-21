package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	from, to, cost int
}

// INF id Infinite Value
const INF = 1000000000

// Edge
var es []edge

// Shortest Distance
var d []int

// V is number of vertexes
var V int

// E is number of edges
var E int

// Solve Shortest Path distance between two points
func bellmanFordMethod(s int) {
	d[s] = 0
	for {
		update := false
		for i := 0; i < 2*E; i++ {
			e := es[i]
			if d[e.from] != INF && d[e.to] > d[e.from]+e.cost {
				d[e.to] = d[e.from] + e.cost
				update = true
			}
		}
		if !update {
			break
		}
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get number of vertexes
	sc.Scan()
	V, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get number of edges
	sc.Scan()
	E, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get Edge info
	es = []edge{}
	for i := 0; i < E; i++ {
		sc.Scan()
		txtedgeinfo := strings.Split(strings.TrimSpace(sc.Text()), " ")
		_from, _ := strconv.Atoi(txtedgeinfo[0])
		_to, _ := strconv.Atoi(txtedgeinfo[1])
		_cost, _ := strconv.Atoi(txtedgeinfo[2])
		es = append(es, edge{from: _from, to: _to, cost: _cost})
		es = append(es, edge{from: _to, to: _from, cost: _cost})
	}

	// Initialize shortest distance array
	d = []int{}
	for i := 0; i < V; i++ {
		d = append(d, INF)
	}

	// Calc shortest path length between start to end
	bellmanFordMethod(0)

	for i, distance := range d {
		fmt.Printf("%d : %d\n", i, distance)
	}
	// Show result
	fmt.Printf("Result :%d\n", d[V-1])

}
