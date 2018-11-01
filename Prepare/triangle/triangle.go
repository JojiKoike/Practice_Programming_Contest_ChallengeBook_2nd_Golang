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
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	parts := strings.Split(strings.TrimSpace(scanner.Text()), ",")
	intparts := []int{}
	for _, part := range parts {
		a, _ := strconv.Atoi(part)
		intparts = append(intparts, a)
	}

	var ans = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				var len = intparts[i] + intparts[j] + intparts[k]
				var longest = math.Max(float64(intparts[i]),
					math.Max(float64(intparts[j]), float64(intparts[k])))
				if int(longest) < len-int(longest) {
					fans := math.Max(float64(ans), float64(len))
					ans = int(fans)
				}
			}
		}
	}
	fmt.Println(ans)
}
