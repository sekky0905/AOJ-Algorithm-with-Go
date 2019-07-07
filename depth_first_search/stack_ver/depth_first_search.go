package main

import (
	"bufio"
	"bytes"
	"errors"
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
	value int
	color
	foundTime     int
	completedTime int
}

var nodes []node

type simpleStack []int

func (s simpleStack) push(v int) {
	_ = append(s, v)
}

func (s simpleStack) pop() (int, error) {
	length := len(s)
	if length < 1 {
		return -1, errors.New("index out of range")
	}
	stack[len(stack)-1] = 0
}

func (s simpleStack) top() (int, error) {
	length := len(s)
	if length < 1 {
		return -1, errors.New("index out of range")
	}
	return s[length-1], nil
}

var stack simpleStack

// deepFirstSearch は深さ優先探索を行う。
func deepFirstSearch() error {
	nodes = make([]node, n, n)
	for i := range nodes {
		nodes[i].color = white
	}

	for i := 0; i < n; i++ {
		if nodes[i].color == white { // 未訪問の頂点を深さ優先探索する
			if err := deepFirstSearchVisit(i); err != nil {
				return err
			}
		}
	}
	return nil
}

func print() {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", i+1, nodes[i].foundTime, nodes[i].completedTime))
	}
	fmt.Print(buf.String())
}

// deepFirstSearchVisit は、深さ優先探索で訪問する。
func deepFirstSearchVisit(u int) error {
	stack.push(u) // 訪問する頂点uをスタックに追加

	// 頂点uを訪問する
	timeCounter++
	nodes[u].color = gray // 訪問したのでuのcolorをgray(訪問済みにする)
	nodes[u].foundTime = timeCounter

	for len(stack) != 0 {
		u, err := stack.pop()
		if err != nil {
			return err
		}

		// 隣接頂点を取得する
		v := getNextNode(u)
		if v != -1 { // 隣接頂点が見つかったら
			if nodes[v].color == white { // 番号順に取得した頂点を訪問する
				timeCounter++
				nodes[v].color = gray
				nodes[v].foundTime++
				stack.push(v) // 頂点がvに移動したので、stackに追加
			}
		} else { // 隣接頂点が見つからなかったら
			stack.pop()
			timeCounter++
			nodes[u].color = black // 隣接頂点が見つからなかったということは、今回の頂点の探索は完了したということ
			nodes[v].completedTime++
		}
	}
	return nil
}

// getNextNode は隣接する頂点を番号順に取得する。
func getNextNode(u int) int {
	for v := nodes[u].value; v < n; v++ {
		nodes[u].value = v + 1
		if adjacentMatrix[u][v] { // uが対象の値で、vがuの今回の隣接頂点
			return v
		}
	}
	return -1
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
	if err := deepFirstSearch(); err != nil {
		panic(err)
	}
	print()
}
