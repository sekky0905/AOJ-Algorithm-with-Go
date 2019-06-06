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

const empty = -1

// setDegree は、二分木のnodeの深さを設定する。
func setDegree(index, degree int) {
	if tree[index].id == empty {
		return
	}

	tree[index].degree = degree
	// 左子node
	setDegree(tree[index].left, degree+1)
	// 右子node
	setDegree(tree[index].right, degree+1)
}

// setHeight は、二分木のnodeの高さを設定する。
func setHeight(index int) int {
	currentNode := tree[index]
	leftHeight, rightHeight := 0, 0
	if currentNode.left != empty {
		leftHeight = setHeight(currentNode.left) + 1 // +1 は、root分
	}

	if currentNode.right != empty {
		rightHeight = setHeight(currentNode.right) + 1
	}

	tree[index].height = max(leftHeight, rightHeight)
	return tree[index].height
}

func max(a, b int) int {
	if a == b {
		panic("a equals b!")
	}
	if a > b {
		return a
	}
	return b
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
