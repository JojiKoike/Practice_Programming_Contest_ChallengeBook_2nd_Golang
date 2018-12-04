package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Gravity Acceleration (m/s^2)
const g = 10.0

// N is number of balls
var N int

// H is height
var H int

// R is ball radius
var R int

// T is elapsed time
var T int

// Ball Final Position
var y []float64

// Calclate Ball position y-coordinate value
func calc(T int) float64 {
	if T < 0 {
		return float64(H)
	}
	t := math.Sqrt(float64(2*H) / g)
	k := int(float64(T) / t)
	var d float64
	if k%2 == 0 {
		d = float64(T) - float64(k)*t
	} else {
		d = float64(k)*t + t - float64(T)
	}
	return float64(H) - g*d*d/2.0
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	//  Get N H R T
	sc.Scan()
	nhrt := strings.Split(sc.Text(), " ")
	N, _ = strconv.Atoi(nhrt[0])
	H, _ = strconv.Atoi(nhrt[1])
	R, _ = strconv.Atoi(nhrt[2])
	T, _ = strconv.Atoi(nhrt[3])

	// Calculation
	y = make([]float64, N)
	for i := 0; i < N; i++ {
		y[i] = calc(T - i)
	}
	sort.Float64s(y)

	// Show Result
	for i := 0; i < N; i++ {
		fmt.Printf("%.2f", y[i]+float64(2*R*i)/100.0)
		if i+1 == N {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}
