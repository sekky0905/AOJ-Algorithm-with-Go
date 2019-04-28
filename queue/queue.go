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
	maxQueue = 100005
	sep      = " "
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// 最初の数字2つ
	s := nextLine()
	split := strings.Split(s, sep)

	n, err := getNum(split)
	if err != nil {
		panic(err)
	}

	num, qms := n[0], n[1]

	// プロセス
	p, err := getQueue(num)
	if err != nil {
		panic(err)
	}

	// CPUをシミュレートする
	if err := Simulate(p, qms); err != nil {
		panic(err)
	}
}

// nextLine は、次の文字列を返す。
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// getNum は、[]stringからintを返す
func getNum(s []string) ([]int, error) {
	list := make([]int, 2, 2)

	for i, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		list[i] = n
	}
	return list, nil
}

// getQueue は、Processの元となるデータを標準入力から取得して、Processに変換する。
func getQueue(n int) (*Queue, error) {
	list := [maxQueue]*Process{}
	q := &Queue{
		head: 0,
		tail: 0,
		list: list,
	}

	for i := 0; i < n; i++ {
		s := nextLine()

		split := strings.Split(s, sep)

		pt, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		p := &Process{
			Name:               split[0],
			RemainsProcessTime: pt,
		}

		if err := q.EnQueue(p); err != nil {
			return nil, err
		}
	}
	return q, nil
}

// Simulate は、CPUの動作をシミュレートする。
func Simulate(q *Queue, qms int) error {
	total := 0

	for q.head == q.head {
		// Queueの中から、1つProcessを取り除く
		p, err := q.DeQueue()
		if err != nil {
			return err
		}

		// RemainsProcessTimeの残り時間か、qmsの小さい方を引く
		x := min(p.RemainsProcessTime, qms)

		// 経過時間に追加する
		total += x

		// 計算する
		p.RemainsProcessTime -= x

		// 残りがまだ存在すれば、Queueに格納する
		if p.RemainsProcessTime != 0 {
			if err := q.EnQueue(p); err != nil {
				return err
			}
		} else {
			fmt.Printf("%s %d\n", p.Name, total)
		}

		// Queueに残りがなくなったら、終了
		if q.IsEmpty() {
			break
		}
	}
	return nil
}

// min は、aとbを比較して小さい方を返す。
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Process は、CPUのプロセスを表す。
type Process struct {
	Name               string
	RemainsProcessTime int //
}

// Queue は、Queueを表す。
// リングバッファとして表現する。
// 空と満杯を区別するため、空の場合はhead=tailとし、満杯の場合は(tail+1=head)%Maxとする
type Queue struct {
	head int                // 配列内の要素が格納されている最初の位置
	tail int                // 配列内の要素が格納されている最後の位置
	list [maxQueue]*Process // Queueの配列
}

// EnQueue は、キューの最後に要素を格納する。
func (q *Queue) EnQueue(el *Process) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.list[q.tail] = el
	if q.tail+1 == maxQueue {
		q.tail = 0 // 先頭に戻す
		return nil
	}

	q.tail++
	return nil
}

// EnQueue は、キューの先頭から要素を取得する。
func (q *Queue) DeQueue() (*Process, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	el := q.list[q.head] // FIFO
	if q.head+1 == maxQueue {
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
	return q.head == (q.tail+1)%maxQueue
}
