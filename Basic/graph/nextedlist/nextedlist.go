package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get Number of vertexes
	sc.Scan()
	V, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get Number of edges
	sc.Scan()
	E, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))

	// Get connection info and build nexted list
	G := make([][]int, V)
	for i := 0; i < E; i++ {
		sc.Scan()
		txtconnectioninput := strings.Split(strings.TrimSpace(sc.Text()), " ")
		s, _ := strconv.Atoi(txtconnectioninput[0])
		t, _ := strconv.Atoi(txtconnectioninput[1])
		G[s] = append(G[s], t)
		G[t] = append(G[t], s)
	}

	for _, g := range G {
		fmt.Printf("%d ", g)
	}
	fmt.Print("\n")
}
