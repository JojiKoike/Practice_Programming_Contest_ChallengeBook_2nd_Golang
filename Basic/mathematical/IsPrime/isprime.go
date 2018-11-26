package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n != 1
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	if isPrime(n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
