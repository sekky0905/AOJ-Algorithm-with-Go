package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type color string

const (
	white    color = "WHITE" // 訪問前のを表す
	gray     color = "GRAY"  // 訪問したことを表す
	black    color = "BLACK" // 完了を表す
	infinity       = 1000000000
)

var (
	// adjacentMatrix は、隣接リストを表す。
	adjacentMatrix [][]node
	edges          []edge
	n              int
)

// node は、頂点を表す。
type node struct {
	key    int
	weight int
}

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

// edge は、頂点を表す。
type edge struct {
	distance int // 最短コストを表す。
	color        // 訪問状態を表す。
}

// initEdges は、edgesを初期化する。
func initEdges() {
	edges = make([]edge, n, n)
	for i := range edges {
		edges[i].color = white
		edges[i].distance = infinity
	}
}

// dijkstra は、ダイクストラのアルゴリズムを表す。
func dijkstra() {
	pq := newPriorityQueue()
	initEdges()

	// 0(始点)に訪問する。
	edges[0].distance = 0
	elm := &element{
		key:      0,
		priority: 0,
	}
	pq.push(elm)
	edges[0].color = gray

	for len(pq) != 0 {
		f := pq.pop()
		u := f.key
		edges[u].color = black // 訪問したことにする

		// 現在の最小値を取り出して、それが最短でなければ今回はスキップする
		if edges[u].distance < f.priority*-1 {
			continue
		}

		for j := 0; j < len(adjacentMatrix[u]); j++ { // リスト[index]を回る
			v := adjacentMatrix[u][j].key
			if edges[v].color == black {
				continue
			}

			if edges[v].distance > edges[u].distance+adjacentMatrix[u][j].weight {
				edges[v].distance = edges[u].distance + adjacentMatrix[u][j].weight
			}

			pq.push(&element{
				key:      v,
				priority: edges[v].distance * -1,
			})
			edges[v].color = gray
		}
	}
}

func print() {
	var buf bytes.Buffer
	for i, v := range edges {
		buf.WriteString(fmt.Sprintf("%d %d\n", i, v.distance))
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
