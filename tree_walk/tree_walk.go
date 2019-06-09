package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	parent int
	left   int
	right  int
}

var tree []node

const (
	empty = -1
)

// preParse は、root→左部分木→右部分木の順番で接点を表示する。
func preParse(index int) {
	if index == empty {
		return
	}

	fmt.Printf(" %d", index)
	// 左部分木
	preParse(tree[index].left)
	// 右部分木
	preParse(tree[index].right)
}

// inParse は、左部分木→root→右部分木の順番で接点を表示する。
func inParse(index int) {
	if index == empty {
		return
	}

	inParse(tree[index].left)
	fmt.Printf(" %d", index)
	inParse(tree[index].right)
}

// 左部分木→右部分木→rootの順番で接点を表示する。
func postParse(index int) {
	if index == empty {
		return
	}
	postParse(tree[index].left)
	postParse(tree[index].right)
	fmt.Printf(" %d", index)
}

type orderType string

const (
	preOrder  orderType = "Preorder"
	inOrder   orderType = "Inorder"
	postOrder orderType = "Postorder"
)

var sc = bufio.NewScanner(os.Stdin)

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func setParent(index, id int) {
	if index != empty {
		tree[index].parent = id
	}
}

func initTree(n int) {
	tree = make([]node, n, n)
	for i := 0; i < n; i++ {
		tree[i] = node{
			parent: empty,
		}
	}
}

func print(root int) {
	fmt.Printf("%s\n", preOrder)
	preParse(root)
	fmt.Println()

	fmt.Printf("%s\n", inOrder)
	inParse(root)
	fmt.Println()

	fmt.Printf("%s\n", postOrder)
	postParse(root)
	fmt.Println()
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanToInt()
	initTree(n)

	for i := 0; i < n; i++ {
		id, left, right := scanToInt(), scanToInt(), scanToInt()
		tree[id].left = left
		tree[id].right = right

		setParent(left, id)
		setParent(right, id)

	}

	root := 0
	for i := 0; i < n; i++ {
		if tree[i].parent == empty {
			root = i
		}
	}

	print(root)
}
