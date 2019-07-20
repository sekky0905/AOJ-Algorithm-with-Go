package main

import (
	"bufio"
	"os"
	"strconv"
)

type color string

const (
	white     color = "WHITE" // 訪問前のを表す
	gray      color = "GRAY"  // 訪問したことを表す
	black     color = "BLACK" // 完了を表す
	maxLength       = 100
	infinity        = -99
)

// node は、頂点を表す。
type node struct {
	distance int // 最短コストを表す。
	parent   int // 親のkeyを表す。
	color        // 訪問状態を表す。
}

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]int
	nodes          []node
	n              int
)

// dijkstra は、ダイクストラのアルゴリズムを表す。
func dijkstra() {
	initNodes()

	// 始点sの設定を行う
	nodes[0].distance = 0 // 最小コストは0
	nodes[0].parent = -1  // 始点のためparentは存在しない

	for {
		// 最小コストを記録する
		minCost := infinity
		var u int

		for i := 0; i < n; i++ {
			if nodes[i].color != black && nodes[i].distance < minCost { // 訪問済みではない && ここまで記録してきた最小コストよりも小さい場合
				minCost = nodes[i].distance
				u = i
			}
		}

		// 最小コストの記録が変更されていないということは、訪問先がなかったということなので終了
		if minCost == infinity {
			break
		}

		nodes[u].color = black

		for v := 0; v < n; v++ {
			if nodes[v].color != black && adjacentMatrix[u][v] != infinity { // 訪問済みではない && uとvに辺が存在する
				if nodes[u].distance+adjacentMatrix[u][v] < nodes[v].distance { // uの最小コスト+uとvの辺の重さがvの最小コストよりも小さい場合
					nodes[v].distance = nodes[u].distance + adjacentMatrix[u][v] // vの最小コストを入れ替え
					nodes[v].parent = u                                          // u → v ということになるので、vのparentをuにする
					nodes[v].color = gray                                        // vに訪問中
				}
			}
		}
	}
}

// initNodes は、nodesを初期化する。
func initNodes() {
	nodes := make([]node, n, n)
	for i := range nodes {
		nodes[i].color = white
		nodes[i].distance = infinity
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

func main() {
	sc.Split(bufio.ScanWords)

	n = scanToInt()

	for i := 0; i < n; i++ {
		// 頂点の番号
		u := scanToInt()
		// uの出次数
		k := scanToInt()

		for i := 0; i < k; i++ {
			// v = uに隣接する頂点の番号, u - v 間の有向辺の重み
			v, c := scanToInt(), scanToInt()
			adjacentMatrix[u][v] = c
		}
	}

	dijkstra()
}
