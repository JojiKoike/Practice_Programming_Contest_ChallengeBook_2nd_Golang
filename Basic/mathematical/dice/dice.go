package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extgcd(a int, b int, x *int, y *int) int {
	d := a
	if b != 0 {
		d = extgcd(b, a%b, y, x)
		*y -= (a / b) * (*x)
	} else {
		*x = 1
		*y = 0
	}
	return d
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Inputs
	sc.Scan()
	inputtxt := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(inputtxt[0])
	b, _ := strconv.Atoi(inputtxt[1])

	x := 1
	y := 1
	fmt.Printf("d: %d\n", extgcd(a, b, &x, &y))
	fmt.Printf("Result : %d, %d\n", x, y)
}
