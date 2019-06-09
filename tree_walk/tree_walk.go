package main

type node struct {
	id    int
	left  int
	right int
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

	print("%d ", index)
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
	print("%d ", index)
	inParse(tree[index].right)
}

// 左部分木→右部分木→rootの順番で接点を表示する。
func postParse(index int) {
	if index == empty {
		return
	}
	postParse(tree[index].left)
	postParse(tree[index].right)
	print("%d ", index)
}

type orderType string

const (
	preOrder  orderType = "Preorder"
	inOrder   orderType = "Inorder"
	postOrder orderType = "Postorder"
)

func main() {

}
