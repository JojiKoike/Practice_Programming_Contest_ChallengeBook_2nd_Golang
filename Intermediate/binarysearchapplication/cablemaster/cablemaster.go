package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// INF is infinite value
const INF = 100000.0

// EPS is acceptable error amount
const EPS = 1.0E-10

// N is number of cables
var N int

// K is number of same length cables
var K int

// L is the length list of given cables
var L []float64

func judge(x float64) bool {
	num := 0
	for i := 0; i < N; i++ {
		num += int(L[i] / x)
	}
	return num >= K
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get N, K
	sc.Scan()
	NK := strings.Split(strings.TrimSpace(sc.Text()), " ")
	N, _ = strconv.Atoi(NK[0])
	K, _ = strconv.Atoi(NK[1])

	// GET L[]
	sc.Scan()
	linputs := strings.Split(strings.TrimSpace(sc.Text()), " ")
	L = []float64{}
	for _, linput := range linputs {
		finput, _ := strconv.ParseFloat(linput, 64)
		L = append(L, finput)
	}

	// Solve
	var lb = 0.0
	var ub = INF
	for ub-lb > EPS {
		mid := (lb + ub) / 2.0
		if judge(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}

	fmt.Printf("%.2f\n", math.Floor(ub*100)/100)

}
