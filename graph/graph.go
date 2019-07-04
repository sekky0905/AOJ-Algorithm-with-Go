package main

import (
	"bufio"
	"os"
	"strconv"
)

// adjacentMatrix は、隣接行列を表す。
var adjacentMatrix [][]int

// initAdjacentMatrix は、adjacentMatrixを初期化する。
func initAdjacentMatrix(n int) {
	adjacentMatrix = make([][]int, n, n)
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

}
