package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	to, cost int
}

type Pair struct {
	distance int
	node     int
	index    int
}

type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	pair := x.(*Pair)
	pair.index = n
	*pq = append(*pq, pair)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(pair *Pair, distance int, node int) {
	pair.node = node
	pair.distance = distance
	heap.Fix(pq, pair.index)
}

// INF is Infinite Value
const INF = 10000000000

// Edge
var G [][]edge

// Shortest Distance
var d []int

// V is number of vertexes
var V int

// E is nuber of edges
var E int

func dijkstra(s int) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	d[s] = 0
	heap.Push(pq, &Pair{distance: 0, node: s})
	for pq.Len() > 0 {
		p := heap.Pop(pq).(*Pair)
		v := p.node
		if d[v] >= p.distance {
			for i := 0; i < len(G[v]); i++ {
				e := G[v][i]
				if d[e.to] > d[v]+e.cost {
					d[e.to] = d[v] + e.cost
					pair := &Pair{
						distance: d[e.to],
						node:     e.to,
					}
					heap.Push(pq, pair)
				}
			}
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

	// Get Edge Info
	G = make([][]edge, V)
	for i := 0; i < E; i++ {
		sc.Scan()
		txtedgeinfo := strings.Split(strings.TrimSpace(sc.Text()), " ")
		_from, _ := strconv.Atoi(txtedgeinfo[0])
		_to, _ := strconv.Atoi(txtedgeinfo[1])
		_cost, _ := strconv.Atoi(txtedgeinfo[2])
		G[_from] = append(G[_from], edge{to: _to, cost: _cost})
		G[_to] = append(G[_to], edge{to: _from, cost: _cost})
	}

	// Initialize shortes distance array
	d = []int{}
	for i := 0; i < V; i++ {
		d = append(d, INF)
	}

	// Calc shortest path length between start to end
	dijkstra(0)

	fmt.Printf("Result : %d\n", d[V-1])
}
