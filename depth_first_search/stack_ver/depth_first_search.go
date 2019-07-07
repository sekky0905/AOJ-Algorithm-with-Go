package main

import "errors"

const maxLength = 100

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]bool
	// timeCounter は全体の時刻を表す。
	timeCounter int
	n           int
)

type color string

const (
	white color = "WHITE" // 訪問前のを表す
	gray  color = "GRAY"  // 訪問したことを表す
	black color = "BLACK" // 完了を表す
)

// node は、グラフにおける頂点を表す。
type node struct {
	color
	foundTime     int
	completedTime int
}

var nodes []node

type simpleStack []int

func (s simpleStack) push(v int) {
	_ = append(s, v)
}

func (s simpleStack) pop() (int, error) {
	length := len(s)
	if length < 1 {
		return 0, errors.New("index out of range")
	}
	return s[length-1], nil
}

func main() {

}
