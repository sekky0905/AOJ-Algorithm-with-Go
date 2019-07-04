package main

import (
	"bufio"
	"os"
	"strconv"
)

// adjacentMatrix は、隣接行列を表す。
var adjacentMatrix [][]int

var sc = bufio.NewScanner(os.Stdin)

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {}
