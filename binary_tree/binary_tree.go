package main

import (
	"bufio"
	"os"
	"strconv"
)

// node は、節を表す。
type node struct {
	id     int
	parent int // 親のnode
	left   int // 左子のnode
	right  int // 右子のnode
	degree int // 深さ(rootからnodeまでの長さ )
	height int // 高さ(nodeから葉までの長さの最大)
}

var tree []*node

var sc = bufio.NewScanner(os.Stdin)

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	n := scanToInt()
	tree := make([]*node, n, n)

	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		id, left, right := scanToInt(), scanToInt(), scanToInt()
		node := &node{
			id:    id,
			left:  left,
			right: right,
		}

		tree[i] = node
	}

}
