package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	exist = 1
)

// adjacentMatrix は、隣接行列を表す。
var adjacentMatrix [][]int

// initAdjacentMatrix は、adjacentMatrixを初期化する。
func initAdjacentMatrix(n int) {
	adjacentMatrix = make([][]int, n, n)
}

// pushRowToAdjacentMatrix は、行をadjacentMatrixのnum番目に格納する。
func pushRowToAdjacentMatrix(num int, row []int) {
	for _, column := range row {
		adjacentMatrix[num][column] = exist
	}
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

func scanToText() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n := scanToInt()
	initAdjacentMatrix(n)

	for i:= 0; i< n; i++ {
		s := scanToText()
		rowStr := strings.Split(s, " ")
		length := len(rowStr)
		row := make([]int, length, length)
		for i, strV := range rowStr {
			numV, err := strconv.Atoi(strV)
			if err != nil {
				panic(err)
			}
			row[i] = numV
		}
		pushRowToAdjacentMatrix(length, row)
	}
}
