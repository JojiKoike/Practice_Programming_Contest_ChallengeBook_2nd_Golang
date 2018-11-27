package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var prime []int
var isPrime []bool

func sieve(n int) int {
	prime = []int{}
	isPrime = make([]bool, n+1)
	p := 0
	for i := 0; i <= n; i++ {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prime = append(prime, i)
			p++
			for j := 2 * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return p
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// Get n
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	res := sieve(n)
	fmt.Printf("Result: %d\n", res)
	if len(prime) <= 0 {
		fmt.Printf("Detected Primes: %v\n", prime)
	}
}
