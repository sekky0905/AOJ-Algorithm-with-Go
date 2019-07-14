package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

const (
	// maxCost は、頂点→頂点の移動の最大のコストを表す。
	maxCost int16 = math.MaxInt16
	// 訪問ステータス
	// white は、訪問前を表す。
	white color = "BeforeVisit"
	// gray は、訪問中を表す。
	gray color = "VisitingNow"
	// black は、訪問済みを表す。
	black color = "AfterVisit"
	// root は、primで訪問する一番最初の点を表す。
	root = -1
)

type (
	// color は、頂点の訪問状態を色で表す。
	color string
	// node は、minimum spanning treeの要素を表す。
	node struct {
		color
		minimumWeight int16
		parent        int16
	}
)

var (
	// n は、頂点の個数を表す。
	n int
	// minimumSpanningTree は、最小全域木を表す。
	minimumSpanningTree []node
	// adjacentMatrix は、隣接行列を表す。
	// 隣接する2頂点の最小コストを格納する。
	// iとjの最小コストは、以下。
	// adjacentMatrix[i][j] = 最小コスト
	// adjacentMatrix[j][i] = 最小コスト
	adjacentMatrix [][]int16
)

// initMinimumSpanningTree は、minimumSpanningTreeを初期化する。
func initMinimumSpanningTree() {
	minimumSpanningTree = make([]node, n, n)

	for i := range minimumSpanningTree {
		minimumSpanningTree[i].color = white
		minimumSpanningTree[i].minimumWeight = maxCost
	}
}

// prim は、primのアルゴリズムを表す。
func prim() int16 {
	// 最小全域木を初期化する。
	initMinimumSpanningTree()

	// 始点となる頂点の設定を行う。
	minimumSpanningTree[0].minimumWeight = 0
	minimumSpanningTree[0].parent = root

	var sum int16 = 0

	for {
		var u int16 = root
		minCost := maxCost
		for i := 0; i < n; i++ {
			// 訪問完了前かつ現時点での最小コストよりコストが小さい場合
			if minimumSpanningTree[i].color != black && minimumSpanningTree[i].minimumWeight < minCost {
				minCost = minimumSpanningTree[i].minimumWeight
				u = int16(i)
			}
		}

		// minCostがmaxCostのままということはどこも訪問していない、つまり、この頂点以下に結ばれている頂点がないことになるので、loopを終了
		if minCost == maxCost {
			break
		}

		// コストが最小の頂点を訪問したので、訪問済みにする。
		minimumSpanningTree[u].color = black

		for v := 0; v < n; v++ {
			// 訪問前かつ、最小コストに変更があった場合
			if minimumSpanningTree[v].color != black && adjacentMatrix[u][v] != maxCost {
				if minimumSpanningTree[v].minimumWeight > adjacentMatrix[u][v] {
					minimumSpanningTree[v].minimumWeight = adjacentMatrix[u][v]
					minimumSpanningTree[v].parent = u
					minimumSpanningTree[v].color = gray
				}
			}
		}

		for i := 0; i < n; i++ {
			if minimumSpanningTree[i].parent != root {
				parent := minimumSpanningTree[i].parent
				sum += adjacentMatrix[i][parent]
			}
		}
	}
	return sum
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
