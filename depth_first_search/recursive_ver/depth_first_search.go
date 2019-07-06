package recursive_ver

import "fmt"

const maxLength = 100

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]bool
	// timeCounter は全体の時刻を表す。
	timeCounter int = 0
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

// deepFirstSearch は、深さ優先探索で訪問する。
func deepFirstSearchVisit(n, u int) {
	// 今回訪問した
	timeCounter++
	nodes[u] = node{
		color:     gray,
		foundTime: timeCounter,
	}

	for i := 0; i < n; i++ {
		if !adjacentMatrix[u][i] { // 隣接行列の今回の箇所に存在しない場合
			continue
		}
		if nodes[i].color == white { // 未踏なら再帰的にその頂点から深さ優先探索で訪問する
			deepFirstSearchVisit(n, i)
		}
	}
	// 完了した
	timeCounter++
	nodes[u] = node{
		color:     black,
		foundTime: timeCounter,
	}
}

// deepFirstSearch は、深さ優先探索を行う。
func deepFirstSearch(n int) {
	for _, node := range nodes { // 初期化する
		node.color = white
	}

	for u := 0; u < n; u++ {
		if nodes[u].color == white {
			deepFirstSearchVisit(n, u)
		}
	}
}

func print(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d %d %d\n", i+1, nodes[i].foundTime, nodes[i].completedTime)
	}
}

func main() {

}
