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
	fmt.Println(fact(n))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
