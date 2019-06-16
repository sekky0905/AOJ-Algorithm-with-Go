package main

import (
	"bufio"
	"os"
	"strconv"
)

// node は、節点を表す。
type node struct {
	key    int
	parent *node // 親
	left   *node // 左の子
	right  *node // 右の子
}

// root は、二分探索木のrootを表す。
var root *node

// insert は、treeにnodeを挿入する。
func insert(z *node) {
	var y *node // xの親
	x := root

	for x != nil { // xがnilじゃない、つまり要素が存在する
		y = x // 親を設定
		if z.key < x.key {
			x = x.left // 左の子に進む
		} else {
			x = x.right // 右の子に進む
		}
	}
	// z(今回の挿入するnode)の親をyにする
	z.parent = y

	if y == nil { // yがnilということは、二分探索木は空だということ
		root = z
	} else if z.key < y.key {
		y.left = z // yの左の子にzを設定する
	} else {
		y.right = z // yの右の子にzを設定する
	}
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

}
