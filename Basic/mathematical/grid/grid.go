package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Pair is x-y coordinate values pair
type Pair struct {
	x, y int
}

func makePair(inputs []string) Pair {
	_x, _ := strconv.Atoi(inputs[0])
	_y, _ := strconv.Atoi(inputs[1])
	return Pair{x: _x, y: _y}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// P1
	sc.Scan()
	p1Inputs := strings.Split(sc.Text(), " ")
	p1 := makePair(p1Inputs)

	// P2
	sc.Scan()
	p2Inputs := strings.Split(sc.Text(), " ")
	p2 := makePair(p2Inputs)

	lx := int(math.Abs(float64(p1.x) - float64(p2.x)))
	ly := int(math.Abs(float64(p1.y) - float64(p2.y)))

	fmt.Printf("lx: %d, ly: %d\n", lx, ly)
	fmt.Printf("Result : %d\n", gcd(lx, ly)-1)
}
