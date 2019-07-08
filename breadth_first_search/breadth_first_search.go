package main

type color string

const (
	white color = "WHITE" // 訪問前のを表す
	gray  color = "GRAY"  // 訪問したことを表す
	black color = "BLACK" // 完了を表す
)

// node は、頂点を表す。
type node struct {
	value int
	color
	distance int
}

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]bool
	nodes          []node
)
