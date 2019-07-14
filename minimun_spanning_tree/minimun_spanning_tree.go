package main

import "math"

const (
	// minCost は、頂点→頂点の移動の最小コストを表す。
	minCost int16 = math.MaxInt16
	// 訪問ステータス
	// white は、訪問前を表す。
	white color = "BeforeVisit"
	// gray は、訪問中を表す。
	gray color = "VisitingNow"
	// black は、訪問済みを表す。
	black color = "AfterVisit"
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
	// minimumSpanningTree は、最小全域木を表す。
	minimumSpanningTree []node
)
