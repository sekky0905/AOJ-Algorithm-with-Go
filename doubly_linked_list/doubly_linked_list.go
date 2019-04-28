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
	insertCommand      = "insert"
	deleteCommand = "delete"
	deleteFirstCommand = "deleteFirst"
	deleteLastCommand  = "deleteLast"
)

// Node は、双方向連結リストの各要素を表す。
type Node struct {
	Data int   // Node自体の値
	Prev *Node // 1つ前のNode
	Next *Node // 1つ後のNode
}

// DoublyLinkedList は、双方向連結リストを表す。
type DoublyLinkedList struct {
	dummy *Node // 番兵

}

// NewDoublyLinkedList は、双方向連結リストを初期化する。
func NewDoublyLinkedList() *DoublyLinkedList {
	dll := &DoublyLinkedList{
		dummy: &Node{},
	}

	dll.dummy.Prev = dll.dummy
	dll.dummy.Next = dll.dummy

	return dll
}

// Insert は、引数で与えられた場所にDoublyLinkedListへのNodeを挿入する。
func (dll *DoublyLinkedList) Insert(data int) {
	// 新しいノードを作成
	newNode := &Node{Data: data}

	// 新しいノードのNextを番兵のNext(番兵)にする
	newNode.Next = dll.dummy.Next
	// 番兵のNextのPrevを新しいノードにする。
	dll.dummy.Next.Prev = newNode

	// 番兵のNextを新しいノードにする。
	dll.dummy.Next = newNode
	// 新しいノードのPrevを番兵にする
	newNode.Prev = dll.dummy
}

// Delete は、引数で与えられたdataを持つNodeを1つ削除する。
func (dll *DoublyLinkedList) Delete(data int) {
	node := dll.findNode(data)

	dll.delete(node)
}

// delete は、引数で与えられたNodeを1つ削除する。
func (dll *DoublyLinkedList) delete(node *Node) {
	if node == nil {
		return
	}

	// Nodeの1つPrevのNextをNodeのNextにする。
	node.Prev.Next = node.Next
	// Nodeの1つNextのPrevをNodeのPrevにする。
	node.Next.Prev = node.Prev
	node = nil
}

// findNode は、指定されたNodeを取得する。
func (dll *DoublyLinkedList) findNode(data int) *Node {
	// 番兵を抜かした最初のNode
	now := dll.dummy.Next

	for now.Data != data {
		if now == dll.dummy { // 最初のNodeがdummyの場合、Nodeが存在しないのでnilを返す
			return nil
		}
		// Nodeを1つ進める
		now = now.Next
	}
	return now
}

// DeleteFirst は、双方向リストの頭のNodeを削除する。
func (dll *DoublyLinkedList) DeleteFirst() {
	// 番兵は、一番前であり一番後ろなので
	dll.delete(dll.dummy.Next)
}

// DeleteFirst は、双方向リストの最後のNodeを削除する。
func (dll *DoublyLinkedList) DeleteLast() {
	// 番兵は、一番前であり一番後ろなので
	dll.delete(dll.dummy.Prev)
}

// ExecuteCommand は、commandを実行する。
func (dll *DoublyLinkedList) ExecuteCommand(inputCommand string) error {
	s := strings.Split(inputCommand, " ")
	command := s[0]

	switch command {
	case insertCommand:
		num, err := strconv.Atoi(s[1])
		if err != nil {
			return err
		}
		dll.Insert(num)
		return nil
	case deleteCommand:
		num, err := strconv.Atoi(s[1])
		if err != nil {
			return err
		}
		dll.Delete(num)
		return nil
	case deleteFirstCommand:
		dll.DeleteFirst()
		return nil
	case deleteLastCommand:
		dll.DeleteLast()
		return nil
	default:
		return errors.New("invalid command")
	}
}

// Print は、Nodeを全て表示する。
func (dll *DoublyLinkedList) Print() {
	// 番兵を抜かした最初のNode
	now := dll.dummy.Next

	for now != dll.dummy {
		fmt.Printf("%d ", now.Data)
		now = now.Next
	}
}

// getTarget は、ターゲットとなる操作のSliceを返す。
func getTarget() []string {
	var length int
	fmt.Scan(&length)

	sc := bufio.NewScanner(os.Stdin)

	target := make([]string, length, length)
	for i := range target {
		if sc.Scan() {
			t := sc.Text()
			target[i] = t
		}
	}

	return target
}

func main() {
	targets := getTarget()
	dll := NewDoublyLinkedList()

	for _, target := range targets {
		dll.ExecuteCommand(target)
	}

	dll.Print()
}
