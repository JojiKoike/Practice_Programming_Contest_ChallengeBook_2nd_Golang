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
	// Number of Rows
	scanner.Scan()
	N, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	// Number of Columns
	scanner.Scan()
	M, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	// Reading Field Info
	field := [][]string{}
	for i := 0; i < N; i++ {
		scanner.Scan()
		row := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		field = append(field, row)
	}

	var res = 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if field[i][j] == "W" {
				field = dfs(i, j, N, M, field)
				res++
			}
		}
	}
	fmt.Printf("Num of Lakes : %d\n", res)
}

func dfs(x int, y int, N int, M int, field [][]string) [][]string {
	field[x][y] = "."

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			var nx = x + dx
			var ny = y + dy
			if 0 <= nx && nx < N && 0 <= ny && ny < M && field[nx][ny] == "W" {
				field = dfs(nx, ny, N, M, field)
			}
		}
	}
	return field
}
