package main

import (
	"bufio"
	"os"
	"strconv"
)

// card は、カードを表す。
type card struct {
	suit string
	Num  int
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
	n := scanToInt()

	list := make([]int, n, n)
	for i := 0; i < n; i++ {
		list[i] = scanToInt()
	}
}
