package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	R, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	sX := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	X := []int{}
	for _, x := range sX {
		ix, _ := strconv.Atoi(x)
		X = append(X, ix)
	}
	sort.Ints(X)
	i := 0
	ans := 0
	for i < N {
		s := X[i]
		i++
		for i < N && X[i] <= s+R {
			i++
		}
		p := X[i-1]
		for i < N && X[i] <= p+R {
			i++
		}
		ans++
	}
	fmt.Printf("Points : %d\n", ans)
}
