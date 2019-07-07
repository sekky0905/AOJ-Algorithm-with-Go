package main

import "errors"

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

func main() {

}
