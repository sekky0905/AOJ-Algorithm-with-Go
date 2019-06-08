package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// node は、節を表す。
type node struct {
	parent int // 親のnode
	left   int // 左子のnode
	right  int // 右子のnode
	depth  int
	height int
}

var tree []node

const empty = -1

// setDepth は、二分木のnodeの深さを設定する。
func setDepth(index, depth int) {
	if index == empty {
		return
	}

	tree[index].depth = depth
	// 左子node
	setDepth(tree[index].left, depth+1)
	// 右子node
	setDepth(tree[index].right, depth+1)
}

// setHeight は、二分木のnodeの高さを設定する。
func setHeight(index int) int {
	var leftHeight, rightHeight int
	if tree[index].left != empty {
		leftHeight = setHeight(tree[index].left) + 1 // +1 は、root分
	}

	if tree[index].right != empty {
		rightHeight = setHeight(tree[index].right) + 1
	}

	tree[index].height = max(leftHeight, rightHeight)
	return tree[index].height
}

// getSibling は、nodeの兄弟のnodeを返す。
func getSibling(index int) int {
	parent := tree[index].parent

	if parent == empty {
		return empty
	}

	// 親の左子nodeが自分自身ではなく、かつ、空でもない場合、それが己の兄弟node
	if tree[parent].left != index && tree[parent].left != empty {
		return tree[parent].left
	}
	// 親の右子nodeが自分自身ではなく、かつ、空でもない場合、それが己の兄弟node
	if tree[parent].right != index && tree[parent].right != empty {
		return tree[parent].right
	}

	return empty
}

func max(a, b int) int {
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

func getRoot() int {
	for i, node := range tree {
		if node.parent == empty {
			return i
		}
	}
	return 0
}

func print(index int) {
	degree := 0
	if tree[index].left != empty {
		degree++
	}

	if tree[index].right != empty {
		degree++
	}

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("node %d: ", index))
	buf.WriteString(fmt.Sprintf("parent = %d, ", tree[index].parent))
	buf.WriteString(fmt.Sprintf("sibling = %d, ", getSibling(index)))
	buf.WriteString(fmt.Sprintf("degree = %d, ", degree))
	buf.WriteString(fmt.Sprintf("depth = %d, ", tree[index].depth))
	buf.WriteString(fmt.Sprintf("height = %d, ", tree[index].height))
	buf.WriteString(fmt.Sprintf("%s\n", getNodeType(index).String()))

	fmt.Print(buf.String())
}

type nodeType string

func (n nodeType) String() string {
	return string(n)
}

const (
	root         nodeType = "root"
	leaf         nodeType = "leaf"
	internalNode nodeType = "internal node"
)

func getNodeType(index int) nodeType {
	if tree[index].parent == empty {
		return root
	}

	if tree[index].left == empty && tree[index].right == empty {
		return leaf
	}

	return internalNode
}

func initTree(n int) {
	tree = make([]node, n, n)
	for i := 0; i < n; i++ {
		tree[i] = node{
			parent: empty,
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := scanToInt()

	initTree(n)
	for i := 0; i < n; i++ {
		index, left, right := scanToInt(), scanToInt(), scanToInt()

		tree[index].left = left
		tree[index].right = right

		// 左子nodeの親nodeはid
		if left != empty {
			tree[left].parent = index
		}

		// 右子nodeの親nodeはid
		if right != empty {
			tree[right].parent = index
		}
	}

	root := getRoot()

	setHeight(root)
	setDepth(root, 0)

	for i := 0; i < n; i++ {
		print(i)
	}
}
