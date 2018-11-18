package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StationHeap []int

func (h StationHeap) Len() int {
	return len(h)
}

func (h StationHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h StationHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StationHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *StationHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// ----- Step 1 . Get inputted values ----- //
	// Get N : Num, of Fuel Supply Stations
	sc.Scan()
	N, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
	// Get L : Distance between start to end
	sc.Scan()
	L, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
	// Get P : Initial Fuel supplied amount
	sc.Scan()
	P, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
	// Get Fuel Stations info : distance from start point
	// and available supply amount
	A := []int{}
	B := []int{}
	for i := 0; i < N; i++ {
		sc.Scan()
		inputinfo := strings.Split(strings.TrimSpace(sc.Text()), " ")
		distance, _ := strconv.Atoi(inputinfo[0])
		amount, _ := strconv.Atoi(inputinfo[1])
		A = append(A, distance)
		B = append(B, amount)
	}
	A = append(A, L)
	B = append(B, 0)

	// Define Priority Queue
	pq := &StationHeap{0}
	heap.Init(pq)
	var ans = 0
	var pos = 0
	var tank = P

	for i := 0; i < N; i++ {
		// Distance between current posotion and next station
		var dx = A[i] - pos

		for tank-dx < 0 {
			if pq.Len() == 0 {
				ans = -1
				break
			}
			tank += heap.Pop(pq).(int)
			ans++
		}

		tank -= dx
		pos = A[i]
		heap.Push(pq, B[i])
	}

	fmt.Printf("Result : %d\n", ans)

}
