package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pieceheap []int

func (h pieceheap) Len() int {
	return len(h)
}

func (h pieceheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h pieceheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pieceheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *pieceheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get Number of Pirces
	sc.Scan()
	N, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
	fmt.Printf("Cut %d pieces\n", N)

	// Get Each piece lentgh and Prepare Priority Queue
	sc.Scan()
	txtlengths := strings.Split(strings.TrimSpace(sc.Text()), " ")
	L := []int{}
	pq := &pieceheap{}
	heap.Init(pq)
	for _, txtlength := range txtlengths {
		length, _ := strconv.Atoi(txtlength)
		L = append(L, length)
		heap.Push(pq, length)
	}

	// Solve
	var ans = 0
	for pq.Len() > 1 {
		// l1 <= l2
		l1 := heap.Pop(pq).(int)
		l2 := heap.Pop(pq).(int)

		fmt.Printf("l1=%d \t l2=%d\n", l1, l2)

		ans += l1 + l2
		heap.Push(pq, l1+l2)
	}

	// Show Result
	fmt.Printf("Result: %d\n", ans)
}
