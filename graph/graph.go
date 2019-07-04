package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	exist = 1
	k = 1
)

// adjacentMatrix は、隣接行列を表す。
var adjacentMatrix [][]int

// initAdjacentMatrix は、adjacentMatrixを初期化する。
func initAdjacentMatrix(n int) {
	adjacentMatrix = make([][]int, n, n)
	for i := range adjacentMatrix {
		adjacentMatrix[i] = make([]int, n)
	}
}

// pushRowToAdjacentMatrix は、行をadjacentMatrixのnum番目に格納する。
func pushRowToAdjacentMatrix(num int, row []int) {
	for _, column := range row {
		adjacentMatrix[num][column] = exist
	}
}

func print() {
	var buf bytes.Buffer
	for _, row := range adjacentMatrix {
		for j, v := range row {
			if j == len(row) -1 {
				buf.WriteString(fmt.Sprintf("%d\n", v))
			} else {
				buf.WriteString(fmt.Sprintf("%d ", v))
			}
		}
	}
	fmt.Print(buf.String())
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
		k, err := strconv.Atoi(rowStr[k])
		if err != nil {
			panic(err)
		}

		row := make([]int, k, k)
		for j, strV := range rowStr[2:] {
			numV, err := strconv.Atoi(strV)
			if err != nil {
				panic(err)
			}
			row[j] = numV -1 // indexのため
		}
		pushRowToAdjacentMatrix(i, row)
	}
	print()
}
