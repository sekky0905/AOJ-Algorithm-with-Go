package main

import "errors"

type color string

const (
	white     color = "WHITE" // 訪問前のを表す
	gray      color = "GRAY"  // 訪問したことを表す
	black     color = "BLACK" // 完了を表す
	maxLength       = 100
)

// node は、頂点を表す。
type node struct {
	value int
	color
	distance int
}

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]bool
	nodes          []node
)

// Queue は、Queueを表す。
// リングバッファとして表現する。
// 空と満杯を区別するため、空の場合はhead=tailとし、満杯の場合は(tail+1=head)%Maxとする
type Queue struct {
	head int // 配列内の要素が格納されている最初の位置
	tail int // 配列内の要素が格納されている最後の位置
	list [maxLength]int
}

// EnQueue は、キューの最後に要素を格納する。
func (q *Queue) EnQueue(el int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.list[q.tail] = el
	if q.tail+1 == maxLength {
		q.tail = 0 // 先頭に戻す
		return nil
	}

	q.tail++
	return nil
}

// EnQueue は、キューの先頭から要素を取得する。
func (q *Queue) DeQueue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty")
	}

	el := q.list[q.head] // FIFO
	if q.head+1 == maxLength {
		q.head = 0
		return el, nil
	}

	q.head++
	return el, nil
}

// IsEmpty は、Queueが空かどうかを確認する。
func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

// IsFull は、Queueが満杯かどうかを確認する。
func (q *Queue) IsFull() bool {
	// 1週回るときに、tailとheadは1つ空けるため
	// また、maxQueueで割るのは、2周目(200000)などを排除するため
	return q.head == (q.tail+1)%maxLength
}