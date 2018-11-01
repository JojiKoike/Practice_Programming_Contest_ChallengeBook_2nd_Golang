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
	m, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	cards := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	intcards := []int{}
	for _, card := range cards {
		intcard, _ := strconv.Atoi(card)
		intcards = append(intcards, intcard)
	}

	var f = false
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					if intcards[i]+intcards[j]+intcards[k]+intcards[l] == m {
						f = true
					}
				}
			}
		}
	}

	if f {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
