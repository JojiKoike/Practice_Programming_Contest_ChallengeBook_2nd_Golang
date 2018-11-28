package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// N is number of Cows house
var N int

// M is number of Agressive Cows
var M int

// x is Cows house corrdinate value list
var x []int

// INF is infinite value
const INF = 100000

func judge(d int) bool {
	last := 0
	for i := 1; i < M; i++ {
		crt := last + 1
		for crt < N && x[crt]-x[last] < d {
			crt++
		}
		if crt == N {
			return false
		}
		last = crt
	}
	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get N, M
	sc.Scan()
	NM := strings.Split(strings.TrimSpace(sc.Text()), " ")
	N, _ = strconv.Atoi(NM[0])
	M, _ = strconv.Atoi(NM[1])

	// Get x
	sc.Scan()
	xinputs := strings.Split(strings.TrimSpace(sc.Text()), " ")
	for _, xinput := range xinputs {
		intxinput, _ := strconv.Atoi(xinput)
		x = append(x, intxinput)
	}
	sort.Ints(x)

	// Solve
	lb := 0
	ub := INF
	for ub-lb > 1 {
		mid := (ub + lb) / 2
		if judge(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}

	// Show Result
	fmt.Printf("Result %d\n", lb)
}
