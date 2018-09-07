package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// オペランドの最大個数
	maxOperandNum = 101
	plus          = "+"
	minus         = "-"
	multiple      = "*"
	separator     = " "
)

func main() {
	// 入力を取得
	target := getTarget()
	// 入力を分割
	strSlice := strings.Split(target, separator)

	var list [maxOperandNum]int
	s := &Stack{
		top:  0,
		list: list,
	}

	for _, v := range strSlice {
		if v == plus || v == minus || v == multiple {
			executeSignalPattern(s, v)
		} else {
			executeOperandPattern(s, v)
		}
	}

	// 最後にstackに残った数字が回答なので、表示する
	answer, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer)
}

// executeSignalPattern は、記号の場合のパターンを実行する。
func executeSignalPattern(s *Stack, signal string) {
	numSlice := make([]int, 2, 2)

	for i := range numSlice {
		num, err := s.Pop()
		if err != nil {
			fmt.Println(err)
		}

		numSlice[i] = num
	}

	// オペランドと記号で計算
	// 取り出すときは、新しいのから出てくるので
	result := calc(numSlice[1], numSlice[0], signal)

	// stackに追加
	if !s.IsFull() {
		if err := s.Push(result); err != nil {
			fmt.Println(err)
		}
	}
}

// executeOperandPattern は、オペランドの場合のパターンを実行する。
func executeOperandPattern(s *Stack, operand string) {
	// 数字に変換してstackに追加
	i, err := strconv.Atoi(operand)
	if err != nil {
		fmt.Println(err)
	}
	if !s.IsFull() {
		if err := s.Push(i); err != nil {
			fmt.Println(err)
		}
	}
}

// getTarget は、ターゲットとなる数式を返す。
func getTarget() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

// Stack は、スタック構造を表す。
type Stack struct {
	top  int // top は、最後の要素が格納されている場所を表す。
	list [maxOperandNum]int
}

// Push は、スタックのトップに引数で与えられた数字を追加する。
func (s *Stack) Push(el int) error {
	if s.IsFull() {
		return errors.New("Stack is Full ")
	}

	s.top++
	s.list[s.top] = el

	return nil
}

// Pop は、スタックのトップの数字を返す。
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack is Empty ")
	}

	s.top--
	return s.list[s.top+1], nil
}

// IsEmpty は、スタックが空かどうかを確認する。
func (s *Stack) IsEmpty() bool {
	// トップが0ならば、何も入っていないため
	return s.top == 0
}

// IsFull は、スタックが満杯かどうかを確認する。
func (s *Stack) IsFull() bool {
	return s.top >= maxOperandNum-1
}

// calc は、計算を行う
func calc(a, b int, signal string) int {
	switch signal {
	case plus:
		return a + b
	case minus:
		return a - b
	default:
		return a * b
	}
}
