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
	txtinputs := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	inputs := []int{}

	// Push
	for _, txtinput := range txtinputs {
		intinput, _ := strconv.Atoi(txtinput)
		inputs = push(inputs, intinput)
	}

	// Pop
	a, slice := pop(inputs)
	for i := 0; i < len(inputs); i++ {
		fmt.Println(a)
		fmt.Println(slice)
		if len(slice) > 0 {
			a, slice = pop(slice)
		}
	}
}

func push(slice []int, elem int) []int {
	slice = append(slice, elem)
	return slice
}

func pop(slice []int) (int, []int) {
	a := slice[len(slice)-1]
	return a, slice[:len(slice)-1]
}
