package main

// Node は、双方向連結リストの各要素を表す。
type Node struct {
	Data int   // Node自体の値
	Prev *Node // 1つ前のNode
	Next *Node // 1つ後のNode
}

// DoublyLinkedList は、双方向連結リストを表す。
type DoublyLinkedList struct {
	dummy Node // 番兵
}

// Insert は、引数で与えられた場所にDoublyLinkedListへのNodeを挿入する。
func (dll *DoublyLinkedList) Insert(num int) {

}

// InitDoublyLinkedList は、双方向連結リストを初期化する。
func InitDoublyLinkedList() {

}

func main() {

}
