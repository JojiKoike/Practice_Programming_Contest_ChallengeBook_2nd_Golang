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
	M, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	field := [][]string{}
	for i := 0; i < N; i++ {
		scanner.Scan()
		row := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		field = append(field, row)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Print(field[i][j] + "\t")
		}
		fmt.Print("\n")
	}
}
