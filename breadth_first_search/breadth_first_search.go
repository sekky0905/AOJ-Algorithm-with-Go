package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color string

const (
	white     color = "WHITE" // 訪問前のを表す
	gray      color = "GRAY"  // 訪問したことを表す
	black     color = "BLACK" // 完了を表す
	maxLength       = 100
	infinity        = -99
)

// node は、頂点を表す。
type node struct {
	color
	distance int
}

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]bool
	nodes          []node
	n              int
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

func initQueue() *Queue {
	list := [maxLength]int{}
	return &Queue{
		head: 0,
		tail: 0,
		list: list,
	}
}

// breadthFirstSearch は、幅優先探索を行う。
func breadthFirstSearch(s int) error {
	nodes = make([]node, n, n)
	for i := range nodes {
		nodes[i].color = white
		nodes[i].distance = infinity
	}

	q := initQueue()

	// 始点を訪問する
	nodes[s].color = gray
	// 始点との距離は0(自分自身なので)
	nodes[s].distance = 0

	if err := q.EnQueue(s); err != nil {
		return err
	}

	for !q.IsEmpty() {
		u, err := q.DeQueue()
		if err != nil {
			return err
		}

		for v := 0; v < n; v++ { // 今回の始点nodeから距離1のループ
			if !adjacentMatrix[u][v] { // 隣接nodeがない場合
				continue
			}

			if nodes[v].distance != infinity { // 頂点にたどり着けない場合
				continue
			}
			nodes[v].distance = nodes[u].distance + 1
			if err := q.EnQueue(v); err != nil {
				return err
			}
		}
		nodes[u].color = black
	}
	return nil
}

func print() {
	var buf bytes.Buffer
	for i, v := range nodes {
		value := i + 1 // 1から始まるので
		if v.distance == infinity {
			buf.WriteString(fmt.Sprintf("%d -1\n", value))
		} else {
			buf.WriteString(fmt.Sprintf("%d %d\n", value, v.distance))
		}

	}
	fmt.Print(buf.String())
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

func scanToText() string {
	sc.Scan()
	return sc.Text()
}

// initAdjacentMatrix は、adjacentMatrixを初期化する。
func initAdjacentMatrix() {
	adjacentMatrix = make([][]bool, n, n)
	for i := range adjacentMatrix {
		adjacentMatrix[i] = make([]bool, n)
	}
}

func main() {
	n = scanToInt()
	initAdjacentMatrix()

	for i := 0; i < n; i++ {
		str := scanToText()
		s := strings.Split(str, " ")

		u, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}

		for _, vStr := range s[2:] {
			v, err := strconv.Atoi(vStr)
			if err != nil {
				panic(err)
			}
			adjacentMatrix[u-1][v-1] = true //indexの分引く
		}
	}

	if err := breadthFirstSearch(0); err != nil {
		fmt.Println(err)
	}

	print()
}
