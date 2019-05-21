package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// partition は、パーティションを使用しSortする。
// パーティションの対象は、p〜rとする。
func partition(list []int, p, r int) int {
	x := list[r]
	i := p - 1

	for j := p; j < r; j++ {
		if list[j] <= x {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[r] = list[r], list[i+1]
	return i + 1
}

var sc = bufio.NewScanner(os.Stdin)

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanToInt()
	list := make([]int, n, n)
	for i := 0; i < n; i++ {
		list[i] = scanToInt()
	}

	r := n - 1
	rv := list[r]
	partition(list, 0, r)

	for i, v := range list {
		if v == rv {
			fmt.Printf("[%d] ", v)
		} else if i == len(list) {
			fmt.Println(v)
		} else {
			fmt.Printf("%d ", v)
		}

	}
}
