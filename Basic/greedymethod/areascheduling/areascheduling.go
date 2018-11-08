package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	start, term int
}

type interval []pair

func (it interval) Len() int      { return len(it) }
func (it interval) Swap(i, j int) { it[i], it[j] = it[j], it[i] }

// Order by term
func (it interval) Less(i, j int) bool { return it[i].term < it[j].term }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	txS := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	scanner.Scan()
	txT := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	itv := []pair{}
	for i := 0; i < N; i++ {
		s, _ := strconv.Atoi(txS[i])
		t, _ := strconv.Atoi(txT[i])
		itv = append(itv, pair{s, t})
	}
	sort.Sort(interval(itv))

	ans := 0
	t := 0
	for i := 0; i < N; i++ {
		if t < itv[i].start {
			ans++
			t = itv[i].term
		}
	}
	fmt.Printf("Works : %d\n", ans)
}
