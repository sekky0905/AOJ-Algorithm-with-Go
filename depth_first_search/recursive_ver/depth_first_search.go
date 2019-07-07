package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxLength = 100

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]bool
	// timeCounter は全体の時刻を表す。
	timeCounter int
	n           int
)

type color string

const (
	white color = "WHITE" // 訪問前のを表す
	gray  color = "GRAY"  // 訪問したことを表す
	black color = "BLACK" // 完了を表す
)

// node は、グラフにおける頂点を表す。
type node struct {
	color
	foundTime     int
	completedTime int
}

var nodes []node

// deepFirstSearchVisit は、深さ優先探索で訪問する。
func deepFirstSearchVisit(u int) {
	// 今回訪問した
	timeCounter++
	nodes[u].color = gray
	nodes[u].foundTime = timeCounter

	for i := 0; i < n; i++ {
		if !adjacentMatrix[u][i] { // 隣接行列の今回の箇所に存在しない場合
			continue
		}
		if nodes[i].color == white { // 未踏なら再帰的にその頂点から深さ優先探索で訪問する
			deepFirstSearchVisit(i)
		}
	}
	// 完了した
	timeCounter++
	nodes[u].color = black
	nodes[u].completedTime = timeCounter
}

// deepFirstSearch は、深さ優先探索を行う。
func deepFirstSearch() {
	nodes = make([]node, maxLength, maxLength)
	for i := range nodes { // 初期化する
		nodes[i].color = white
	}

	for u := 0; u < n; u++ {
		if nodes[u].color == white {
			deepFirstSearchVisit(u)
		}
	}
}

func print() {
	for i := 0; i < n; i++ {
		fmt.Printf("%d %d %d\n", i+1, nodes[i].foundTime, nodes[i].completedTime)
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

// initAdjacentMatrix は、adjacentMatrixを初期化する。
func initAdjacentMatrix() {
	adjacentMatrix = make([][]bool, n, n)
	for i := range adjacentMatrix {
		adjacentMatrix[i] = make([]bool, n)
	}
}

func main() {
	n = scanToInt()
	initAdjacentMatrix()

	for i := 0; i < n; i++ {
		str := scanToText()

		s := strings.Split(str, " ")

		u, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}

		for _, vStr := range s[2:] {
			v, err := strconv.Atoi(vStr)
			if err != nil {
				panic(err)
			}
			adjacentMatrix[u-1][v-1] = true //indexの分引く
		}
	}
	deepFirstSearch()
	print()
}
