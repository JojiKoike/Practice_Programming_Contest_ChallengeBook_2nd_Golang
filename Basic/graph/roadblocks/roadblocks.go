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

// G is graph structure data
var G [][]edge

// N is number of cross point (in graph, node)
var N int

// R is number of Roads (in graph, edge)
var R int

// Shortest Distances
var dist []int

// Second Shortest Distance
var dist2 []int

func solve() {
	pq := &PriorityQueue{}
	heap.Init(pq)
	dist[0] = 0
	heap.Push(pq, &Pair{distance: 0, node: 0})

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*Pair)
		v := p.node
		d := p.distance
		if dist2[v] >= d {
			for i := 0; i < len(G[v]); i++ {
				e := G[v][i]
				d2 := d + e.cost
				if dist[e.to] > d2 {
					dist[e.to], d2 = d2, dist[e.to]
					pair := &Pair{
						distance: dist[e.to],
						node:     e.to,
					}
					heap.Push(pq, pair)
				}
				if dist2[e.to] > d2 && dist[e.to] < d2 {
					dist2[e.to] = d2
					pair := &Pair{
						distance: dist2[e.to],
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

	// Get number of cross points
	sc.Scan()
	N, _ = strconv.Atoi(sc.Text())

	// Get number of Roads
	sc.Scan()
	R, _ = strconv.Atoi(sc.Text())

	// Get Edge Info
	G = make([][]edge, N)
	for i := 0; i < R; i++ {
		sc.Scan()
		txtedgeinfo := strings.Split(sc.Text(), " ")
		_from, _ := strconv.Atoi(txtedgeinfo[0])
		_to, _ := strconv.Atoi(txtedgeinfo[1])
		_cost, _ := strconv.Atoi(txtedgeinfo[2])
		G[_from] = append(G[_from], edge{to: _to, cost: _cost})
		G[_to] = append(G[_to], edge{to: _from, cost: _cost})
	}

	// Initialize distance data table
	dist = make([]int, N)
	dist2 = make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = INF
		dist2[i] = INF
	}

	solve()

	fmt.Printf("Result : %d\n", dist2[N-1])
}
