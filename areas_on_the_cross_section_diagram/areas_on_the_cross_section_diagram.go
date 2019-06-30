package main

import (
	"fmt"
)

func main() {
	// 全体の面積
	grossArea := 0

	// \ の位置のStack
	leftSignalPositionStack := &Stack{}

	// inputを取得する
	input := getInput()

	// 各水たまりの面積
	nodes := &Stack{}

	for nowPosition, signal := range input {
		if string(signal) == `\` { // \ であれば、stackに積む
			leftSignalPositionStack.Push(nowPosition)
		} else if string(signal) == "/" {

			if len(leftSignalPositionStack.list) <= 0 {
				continue
			}

			// leftSignalPositionStackのトップを取り出す
			top := leftSignalPositionStack.Pop().(int)

			// 今回の面積を求める
			area := nowPosition - top

			// 今回の面積を総合面積に追加する
			grossArea += area

			for len(nodes.list) > 0 {
				if top < nodes.list[len(nodes.list)-1].(*Node).Position {
					node := nodes.Pop().(*Node)
					area += node.Area // 新たにできた水たまりの面積を直線までの面積に加える
				} else {
					break
				}
			}
			// 各水たまりに加える
			// 各々水たまりも、\と/の数が合う度に面積を加えていく必要がある
			nodes.Push(&Node{Area: area, Position: top})
		}
	}

	length := len(nodes.list)

	fmt.Println(grossArea)
	fmt.Print(length)

	if length > 0 {
		for _, node := range nodes.list {
			fmt.Printf(" %d", node.(*Node).Area)
		}
	}
	fmt.Println()
}

// Node は、その時点での窪み。
type Node struct {
	Area     int
	Position int
}

// getTarget は、入力を受け取る。
func getInput() string {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}
	return input
}

// Stack は、スタック構造を表す。
type Stack struct {
	list []interface{}
}

// Push は、スタックのトップに引数で与えられた数字を追加する。
func (s *Stack) Push(el interface{}) {
	s.list = append(s.list, el)
}

// Pop は、スタックのトップの数字を返す。
func (s *Stack) Pop() (target interface{}) {
	length := len(s.list)
	target = s.list[length-1]
	s.list = s.list[:length-1]
	return
}
