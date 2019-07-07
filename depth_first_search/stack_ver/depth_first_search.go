package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxLength = 1000

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]bool
	// timeCounter は全体の時刻を表す。
	timeCounter int
	n           int
	nodes       []node
	stack       *Stack
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

// Stack は、スタック構造を表す。
type Stack struct {
	top  int // top は、最後の要素が格納されている場所を表す。
	list [maxLength]int
}

// Push は、スタックのトップに引数で与えられた数字を追加する。
func (s *Stack) Push(el int) {
	s.top++
	s.list[s.top] = el
}

// Pop は、スタックのトップの数字を返す。
func (s *Stack) Pop() int {
	s.top--
	return s.list[s.top+1]
}

// IsEmpty は、スタックが空かどうかを確認する。
func (s *Stack) IsEmpty() bool {
	// トップが0ならば、何も入っていないため
	return s.top == 0
}

// IsFull は、スタックが満杯かどうかを確認する。
func (s *Stack) IsFull() bool {
	return s.top >= maxLength-1
}

// Top は、スタックの一番上を返す。
func (s *Stack) Top() int {
	return s.list[s.top]
}

// deepFirstSearch は深さ優先探索を行う。
func deepFirstSearch() error {
	nodes = make([]node, maxLength, maxLength)
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

// deepFirstSearchVisit は、深さ優先探索で訪問する。
func deepFirstSearchVisit(r int) error {
	if !stack.IsFull() {
		stack.Push(r) // 訪問する頂点uをスタックに追加
	}

	// 頂点uを訪問する
	timeCounter++
	nodes[r].color = gray // 訪問したのでuのcolorをgray(訪問済みにする)
	nodes[r].foundTime = timeCounter

	for stack.top != 0 {
		u := stack.Top()

		// 隣接頂点を取得する
		v := getNextNode(u)
		if v != -1 { // 隣接頂点が見つかったら
			if nodes[v].color == white { // 番号順に取得した頂点を訪問する
				timeCounter++
				nodes[v].color = gray
				nodes[v].foundTime = timeCounter
				if !stack.IsFull() {
					stack.Push(v) // 頂点がvに移動したので、stackに追加
				}
			}
		} else { // 隣接頂点が見つからなかったら
			if !stack.IsEmpty() {
				stack.Pop()
			}
			timeCounter++
			nodes[u].color = black // 隣接頂点が見つからなかったということは、今回の頂点の探索は完了したということ
			nodes[u].completedTime = timeCounter
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

func print() {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", i+1, nodes[i].foundTime, nodes[i].completedTime))
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

// initAdjacentMatrix は、adjacentMatrixを初期化する。
func initAdjacentMatrix() {
	adjacentMatrix = make([][]bool, n, n)
	for i := range adjacentMatrix {
		adjacentMatrix[i] = make([]bool, n)
	}
}

func initStack() {
	var list [maxLength]int
	stack = &Stack{
		top:  0,
		list: list,
	}
}

func main() {
	n = scanToInt()
	initAdjacentMatrix()
	initStack()

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
		fmt.Println(err.Error())
	}
	print()
}
