package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	L, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	//n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	ants := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	// Calculate minimum time
	var minT = 0
	var maxT = 0
	for _, ant := range ants {
		intant, _ := strconv.Atoi(ant)
		minT = int(math.Max(float64(minT), math.Min(float64(intant), float64(L-intant))))
		maxT = int(math.Max(float64(maxT), math.Max(float64(intant), float64(L-intant))))
	}
	fmt.Printf("%d %d\n", minT, maxT)
}
