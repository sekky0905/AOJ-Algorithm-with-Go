package main

type color string

const (
	white    color = "WHITE" // 訪問前のを表す
	gray     color = "GRAY"  // 訪問したことを表す
	black    color = "BLACK" // 完了を表す
	infinity       = 1000000000
)

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]int
	nodes          []node
	n              int
)

// element は、priority queueの要素を表す。
type element struct {
	key      int
	priority int
}

// priorityQueue は、priority queueを表す。
type priorityQueue []*element

// push は、priority queueにelementを追加する。
func (pq *priorityQueue) push(elm *element) {
	n := len(*pq)
	elm.key = n
	*pq = append(*pq, elm)
}

// pop は、priority queueから次のelementを削除する。
func (pq *priorityQueue) pop() *element {
	old := *pq
	n := len(old)
	elm := old[n-1]
	*pq = old[0 : n-1]
	return elm
}

func newPriorityQueue() priorityQueue {
	return make(priorityQueue, n)
}

// node は、頂点を表す。
type node struct {
	distance int // 最短コストを表す。
	color        // 訪問状態を表す。
}
