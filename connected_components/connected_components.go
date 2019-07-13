package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	n     int
	graph [][]int
	// groups は、友達の関係を表す。
	// 友達の関係がある場合には、[index]に同一のidが入る
	groups []int
)

const (
	empty = -1
)

// depthFirstSearch は、深さ優先探索を行う。
func depthFirstSearch(r, id int) error {
	// container/listをstack的な用途で使用する
	stackMock := list.New()
	// 訪問中の頂点をstackに追加
	stackMock.PushFront(r)
	// 始点となる頂点にidを付与する
	groups[r] = id

	for stackMock.Len() != 0 {
		// stackの一番上の値を取り出す(最後に格納された奴)
		u, ok := stackMock.Remove(stackMock.Front()).(int)
		if !ok {
			return errors.New("failed to type assertion")
		}

		// uの隣接部分を順番に訪問していく
		for i := 0; i < len(graph[u]); i++ {
			v := graph[u][i]
			if groups[v] == empty {
				groups[v] = id
				stackMock.PushFront(v)
			}
		}
	}
	return nil
}

func assignColor() error {
	id := 1
	// 初期化
	for i := 0; i < n; i++ {
		groups[i] = empty
	}

	for u := 0; u < n; u++ {
		if groups[u] == empty {
			// ここでは、頂点の訪問毎に異なるidをつけている
			// depthFirstSearch では、uを始点として深さ優先探索を行い、その際に訪問したu以降の頂点は皆uと関係があるものと
			// 考え、同一のidを付与する
			if err := depthFirstSearch(u, id); err != nil {
				return err
			}
			id++
		}
	}
	return nil
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

func print(s, t int) {
	if groups[s] == groups[t] {
		fmt.Println("yes")
		return
	}
	fmt.Println("no")
}

func main() {
	sc.Split(bufio.ScanWords)
	n = scanToInt()
	graph = make([][]int, n, n)
	groups = make([]int, n, n)

	m := scanToInt()

	for i := 0; i < m; i++ {
		s, t := scanToInt(), scanToInt()
		// 友達同士お互いの関係をgraph内のお互いの場所に記録する
		graph[s] = append(graph[s], t)
		graph[t] = append(graph[t], s)
	}

	if err := assignColor(); err != nil {
		panic(err)
	}

	q := scanToInt()
	for i := 0; i < q; i++ {
		s, t := scanToInt(), scanToInt()
		print(s, t)
	}
}
