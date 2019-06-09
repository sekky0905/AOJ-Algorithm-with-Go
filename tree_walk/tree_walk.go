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

// preParse は、root→左部分木→右部分木の順番で接点の番号を巡回する。
func preParse(index int) {
	if index == empty {
		return
	}

	// TODO print

	// 左部分木
	preParse(tree[index].left)
	// 右部分木
	preParse(tree[index].right)
}

func main() {

}
