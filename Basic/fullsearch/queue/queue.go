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

	// Enqueue
	for _, txtinput := range txtinputs {
		iniinput, _ := strconv.Atoi(txtinput)
		inputs = enqueue(inputs, iniinput)
	}

	// Dequeue
	a, slice := dequeue(inputs)
	for i := 0; i < len(inputs); i++ {
		fmt.Printf("Dequeued Value = %d : Dequeued Slice = %v\n", a, slice)
		if len(slice) > 0 {
			a, slice = dequeue(slice)
		}
	}

}

func enqueue(slice []int, elem int) []int {
	slice = append(slice, elem)
	return slice
}

func dequeue(slice []int) (int, []int) {
	a := slice[0]
	return a, slice[1:]
}
