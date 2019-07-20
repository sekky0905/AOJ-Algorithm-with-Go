package main

type color string

const (
	white    color = "WHITE" // 訪問前のを表す
	gray     color = "GRAY"  // 訪問したことを表す
	black    color = "BLACK" // 完了を表す
	infinity       = 1000000000
)

// element は、priority queueの要素を表す。
type element struct {
	key      int
	priority int
}

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]int
	n              int
)
