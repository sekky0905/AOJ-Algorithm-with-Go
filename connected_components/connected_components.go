package main

import (
	"container/list"
	"errors"
)

var (
	n     int
	graph [][]int
	color []int
)

const (
	empty = -1
)

// deepFirstSearch は、深さ優先探索を行う。
func deepFirstSearch(r, c int) error {
	// container/listをstack的な用途で使用する
	stackMock := list.New()
	stackMock.PushFront(r)
	color[r] = c

	for stackMock.Len() != 0 {
		// stackの一番上の値を取り出す(最後に格納された奴)
		front := stackMock.Front()
		u, ok := front.Value.(int)
		if !ok {
			return errors.New("failed to type assertion")
		}
		stackMock.Remove(front)

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
			if err := deepFirstSearch(u, id); err != nil {
				return err
			}
		}
	}
	return nil
}
