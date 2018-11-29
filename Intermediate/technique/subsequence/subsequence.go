package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n int

// S is summation minimum
var S int

var a []int

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get n
	sc.Scan()
	n, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))
	// Get S
	sc.Scan()
	S, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))
	// Get a list
	a = []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		aint, _ := strconv.Atoi(sc.Text())
		a = append(a, aint)
	}

	res := n + 1
	s := 0
	t := 0
	sum := 0
	for {
		for t < n && sum < S {
			sum += a[t]
			t++
		}
		if sum < S {
			break
		}
		if res >= t-s {
			res = t - s
		}
		sum -= a[s]
		s++
	}

	if res > n {
		res = 0
	}

	fmt.Printf("Result : %d\n", res)

}
