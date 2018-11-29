package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var n int
var k int
var w []int
var v []int
var y []float64

// EPS is acceptable error
const EPS = 1.0E-10

func judge(x float64) bool {
	y := []float64{}
	for i := 0; i < n; i++ {
		y = append(y, float64(v[i])-x*float64(w[i]))
	}
	sort.Float64s(y)
	sum := 0.0
	for i := 0; i < k; i++ {
		sum += y[n-i-1]
	}
	return sum >= 0
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get n, k
	sc.Scan()
	nk := strings.Split(strings.TrimSpace(sc.Text()), " ")
	n, _ = strconv.Atoi(nk[0])
	k, _ = strconv.Atoi(nk[1])

	// Get weights, values
	w = []int{}
	v = []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		wv := strings.Split(strings.TrimSpace(sc.Text()), " ")
		iw, _ := strconv.Atoi(wv[0])
		iv, _ := strconv.Atoi(wv[1])
		w = append(w, iw)
		v = append(v, iv)
	}

	// Solve
	lb := 0.0
	ub := 1000000.0
	for ub-lb > EPS {
		mid := (ub + lb) / 2.0
		if judge(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}

	fmt.Printf("Result : %.2f\n", ub)

}
