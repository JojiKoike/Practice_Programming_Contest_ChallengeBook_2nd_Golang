package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	k, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	inputs := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	a := []int{}
	for _, input := range inputs {
		intinput, _ := strconv.Atoi(input)
		a = append(a, intinput)
	}
	if dfs(0, a, n, k, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(i int, a []int, n int, k int, sum int) bool {
	if i == n {
		return sum == k
	}
	if dfs(i+1, a, n, k, sum) {
		return true
	}
	if dfs(i+1, a, n, k, sum+a[i]) {
		return true
	}
	return false
}
