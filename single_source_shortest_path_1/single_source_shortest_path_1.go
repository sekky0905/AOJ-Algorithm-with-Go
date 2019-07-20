package main

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
	key      int // 頂点のkeyを表す。
	distance int // 最短コストを表す。
	parent   int // 親のkeyを表す。
	color        // 訪問状態を表す。
}

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]bool
	nodes          []node
	n              int
)

func main() {

}
