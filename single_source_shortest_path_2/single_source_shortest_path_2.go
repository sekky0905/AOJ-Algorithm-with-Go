package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

const infinity = 1000000000

var (
	// distances は、最短経路の集合を表す。
	distances []int
	// adjacentList は、連接リストを表す。
	adjacentList [][]node
)

type (
	// element は、priorityQueueの要素を表す。
	element struct {
		key      int
		distance int
	}

	// priorityQueue は、priority queueを表す。
	priorityQueue []*element

	// node は、頂点を表す。
	node struct {
		key    int
		weight int
	}
)

// Len は、priorityQueueの長さを返す。
func (pq priorityQueue) Len() int { return len(pq) }

// Less は、iがjの要素よりも小さいかどうかを返す。
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

// Swap は、iとjの要素を交換する。
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push は、priorityQueueに要素を格納する。
func (pq *priorityQueue) Push(x interface{}) {
	elm := x.(*element)
	*pq = append(*pq, elm)
}

// Pop は、priorityQueueから要素を取り除く。
func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	elm := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return elm
}

var sc = bufio.NewScanner(os.Stdin)

// dijkstra は、ダイクストラのアルゴリズムを表す。
func dijkstra(pq priorityQueue) {
	distances[0] = 0
	pq.Push(&element{0, 0})

	for i := 1; i < len(distances); i++ {
		distances[i] = infinity
		pq.Push(&element{i, distances[i]})
	}

	heap.Init(&pq)

	for pq.Len() != 0 {
		elm := heap.Pop(&pq).(*element)
		u := elm.key
		for i := 0; i < len(adjacentList[u]); i++ {
			v := adjacentList[u][i].key
			if elm.distance+adjacentList[u][i].weight < distances[v] {
				distances[v] = elm.distance + adjacentList[u][i].weight
				heap.Push(&pq, &element{v, distances[v]})
			}
		}
	}
}

func print() {
	var buf bytes.Buffer
	for i, v := range distances {
		buf.WriteString(fmt.Sprintf("%d %d\n", i, v))
	}
	fmt.Print(buf.String())
}

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

	n := scanToInt()

	adjacentList = make([][]node, n)
	pq := make(priorityQueue, 0, n)
	distances = make([]int, n)

	for i := 0; i < n; i++ {
		u := scanToInt()
		k := scanToInt()
		adjacentList[u] = make([]node, k)
		for j := 0; j < k; j++ {
			adjacentList[u][j].key = scanToInt()
			adjacentList[u][j].weight = scanToInt()
		}
	}

	dijkstra(pq)
	print()
}
