package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get P (Total Pages)
	sc.Scan()
	P, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get a
	a := make([]int, P)
	for i := 0; i < P; i++ {
		sc.Scan()
		ainput, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
		a[i] = ainput
	}

	// Build Set using feature that map key is distinctive
	allmap := make(map[int]struct{})
	for i := 0; i < P; i++ {
		allmap[a[i]] = struct{}{}
	}
	all := []int{}
	for key := range allmap {
		all = append(all, key)
	}
	var n = len(all)

	// Solve
	var s = 0
	var t = 0
	var num = 0
	count := make(map[int]int) // Content -> Encounters mapping
	var res = P
	for {
		for t < P && num < n {
			if count[a[t]] == 0 {
				num++
			}
			count[a[t]]++
			t++
		}
		if num < n {
			break
		}
		if res > t-s {
			res = t - s
		}
		if count[a[s]] == 1 {
			num--
		}
		s++
	}

	// Show Result
	fmt.Printf("Result: %d\n", res)

}
