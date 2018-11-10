package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// INF defines Infinite Value in this theme
const INF = -10000000

type pair struct {
	x, y int
}

var maze = [][]string{} // Maze Data
var n, m int            // Maze Size
var sx, sy int          // Start point coordinate value
var gx, gy int          // Goal point coordinate value

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

var d = [][]int{}

func bfs() int {

	// Init d Table
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < m; j++ {
			row = append(row, INF)
		}
		d = append(d, row)
	}

	queue := []pair{}
	// Enqueue Start Point and set it's distance from start as zero
	queue = append(queue, pair{x: sx, y: sy})
	d[sx][sy] = 0

	for len(queue) > 0 {
		// Dequeue
		var point = queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []pair{}
		}
		// Goal Check
		if point.x == gx && point.y == gy {
			break
		}

		for i := 0; i < 4; i++ {
			nx := point.x + dx[i]
			ny := point.y + dy[i]

			if 0 <= nx && nx < n && 0 <= ny &&
				ny < m && maze[nx][ny] != "#" &&
				d[nx][ny] == INF {
				queue = append(queue, pair{x: nx, y: ny})
				d[nx][ny] = d[point.x][point.y] + 1
			}
		}
	}
	return d[gx][gy]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Get Maze Size
	scanner.Scan()
	size := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	n, _ = strconv.Atoi(size[0])
	m, _ = strconv.Atoi(size[1])

	// Reading Maze Data
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		row = append(row, "#")
		maze = append(maze, row)
	}

	// Get Start and Goal Coordinate Values
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if maze[i][j] == "S" {
				sx, sy = i, j
			} else if maze[i][j] == "G" {
				gx, gy = i, j
			}
		}
	}

	fmt.Printf("Result : %d\n", bfs())
}
