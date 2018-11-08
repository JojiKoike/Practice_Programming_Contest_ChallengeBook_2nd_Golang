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
	queue := list.New()
	for _, txtInput := range txtInputs {
		intInput, _ := strconv.Atoi(txtInput)
		queue.PushBack(intInput)
		fmt.Printf("Queued Value: %d\tQueue Size: %d\n", intInput, queue.Len())
	}
	for queue.Len() > 0 {
		value := queue.Remove(queue.Front())
		fmt.Printf("DeQueued Value: %d\tQueue Size: %d\n", value, queue.Len())
	}
}
