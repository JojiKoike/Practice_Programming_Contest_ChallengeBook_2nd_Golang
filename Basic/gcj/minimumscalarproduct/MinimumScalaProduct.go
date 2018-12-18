package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	v1txt := strings.Split(sc.Text(), " ")
	v1 := []int{}
	for _, v := range v1txt {
		vint, _ := strconv.Atoi(v)
		v1 = append(v1, vint)
	}
	sort.Ints(v1)

	sc.Scan()
	v2txt := strings.Split(sc.Text(), " ")
	v2 := []int{}
	for _, v := range v2txt {
		vint, _ := strconv.Atoi(v)
		v2 = append(v2, vint)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(v2)))

	res := 0
	for i := 0; i < n; i++ {
		res += v1[i] * v2[i]
	}

	fmt.Println(res)

}
