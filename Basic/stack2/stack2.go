package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	txtInputs := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	stack := list.New()
	for _, txtInput := range txtInputs {
		intInput, _ := strconv.Atoi(txtInput)
		stack.PushFront(intInput)
	}
	for stack.Len() > 0 {
		value := stack.Remove(stack.Front())
		fmt.Printf("Popped Value : %d  Stack Size: %v\n", value, stack.Len())
	}
}
