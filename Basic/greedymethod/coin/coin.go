package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	v := [6]int{1, 5, 10, 50, 100, 500}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	coins := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	icoins := []int{}
	for _, coin := range coins {
		icoin, _ := strconv.Atoi(coin)
		icoins = append(icoins, icoin)
	}
	scanner.Scan()
	A, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	var ans = 0

	for i := 5; i >= 0; i-- {
		var t int
		if A/v[i] < icoins[i] {
			t = A / v[i]
		} else {
			t = icoins[i]
		}
		A -= t * v[i]
		ans += t
	}
	fmt.Printf("Result : %d\n", ans)
}
