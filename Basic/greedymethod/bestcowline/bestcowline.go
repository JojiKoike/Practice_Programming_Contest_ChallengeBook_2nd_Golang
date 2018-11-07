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
	N, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	S := strings.TrimSpace(scanner.Text())
	var a, b = 0, N - 1

	for a <= b {
		var left = false
		for i := 0; a+i <= b; i++ {
			if S[a+i] < S[b-i] {
				left = true
				break
			} else if S[a+i] > S[b-i] {
				left = false
				break
			}
		}
		if left {
			fmt.Printf("%c", S[a])
			a++
		} else {
			fmt.Printf("%c", S[b])
			b--
		}
	}
	fmt.Print("\n")
}
