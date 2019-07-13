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
	color []int
)

const (
	empty = -1
)

// depthFirstSearch は、深さ優先探索を行う。
func depthFirstSearch(r, c int) error {
	// container/listをstack的な用途で使用する
	stackMock := list.New()
	stackMock.PushFront(r)
	color[r] = c

	for stackMock.Len() != 0 {
		// stackの一番上の値を取り出す(最後に格納された奴)
		u, ok := stackMock.Remove(stackMock.Front()).(int)
		if !ok {
			return errors.New("failed to type assertion")
		}

		for i := 0; i < len(graph[u]); i++ {
			v := graph[u][i]
			if color[v] == empty {
				color[v] = c
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
		color[i] = empty
	}

	for u := 0; u < n; u++ {
		if color[u] == empty {
			id++
			if err := depthFirstSearch(u, id); err != nil {
				return err
			}
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
	if color[s] == color[t] {
		fmt.Println("yes")
		return
	}
	fmt.Println("no")
}

func main() {
	sc.Split(bufio.ScanWords)
	n = scanToInt()
	graph = make([][]int, n, n)
	color = make([]int, n, n)

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
