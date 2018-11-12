package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	w, v int
}

var n int // Number of Items

// W is Available Total Weight
var W int
var items = []item{}

func rec(i, j int) int {
	var res int
	if i == n {
		res = 0
	} else if j < items[i].w {
		res = rec(i+1, j)
	} else {
		var rec1 = rec(i+1, j)
		var rec2 = rec(i+1, j-items[i].w) + items[i].v
		if rec1 < rec2 {
			res = rec2
		} else {
			res = rec1
		}
	}
	return res
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get Number of items
	sc.Scan()
	n, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get Items Data
	for i := 0; i < n; i++ {
		sc.Scan()
		txtitem := strings.Split(strings.TrimSpace(sc.Text()), " ")
		weight, _ := strconv.Atoi(txtitem[0])
		value, _ := strconv.Atoi(txtitem[1])
		items = append(items, item{w: weight, v: value})
	}

	// Get Available Total Weight
	sc.Scan()
	W, _ = strconv.Atoi(strings.TrimSpace(sc.Text()))

	fmt.Println(rec(0, 5))

}
