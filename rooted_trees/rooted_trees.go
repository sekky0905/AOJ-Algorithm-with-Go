package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// node は、木の接点もしくは葉を表す。
type node struct {
	id       int
	parent   int
	depth    int
	children []int
}

// setDepth は、depth を設定する。
func setDepth(nodes []*node, currentNodeID, depth int) {
	nodes[currentNodeID].depth = depth
	for i := range nodes[currentNodeID].children {
		// 自分の child に当たる node の depth をそれぞれ再帰的に設定していく
		setDepth(nodes, nodes[currentNodeID].children[i], depth+1)
	}
}

func bufferingNodeInfo(node *node, buf *bytes.Buffer) {
	nodeType := "leaf"
	if node.parent == -1 {
		nodeType = "root"
	} else if len(node.children) > 0 {
		nodeType = "internal node"
	}

	childStr := strings.Replace(fmt.Sprintf("%v", node.children), " ", ", ", -1)
	buf.WriteString(fmt.Sprintf("node %d: parent = %d, depth = %d, %s, %s\n", node.id, node.parent, node.depth, nodeType, childStr))
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

func scanToStr() string {
	sc.Scan()
	return sc.Text()
}

const (
	idIndex = iota
	degreeIndex
)

func main() {
	// 高さ n
	height := scanToInt()
	// 高さ height の tree
	tree := make([]*node, height)

	previousID := -1
	for i := 0; i < height; i++ {
		// 1行ずつ取り出す
		// 次数ごとの情報
		row := scanToStr()
		info := strings.Split(row, " ")

		id, err := strconv.Atoi(info[idIndex])
		if err != nil {
			panic(err)
		}

		degree, err := strconv.Atoi(info[degreeIndex])
		if err != nil {
			panic(err)
		}

		// child node を埋める
		children := make([]int, degree, degree)
		for i := 0; i < degree; i++ {
			if children[i], err = strconv.Atoi(info[i+2]); err != nil { // +2は、id と degree の分
				panic(err)
			}
		}

		tree[i] = &node{
			parent:   previousID,
			id:       id,
			children: children,
		}

		previousID = id
	}

	root := -1
	for i, node := range tree {
		if node.parent == -1 {
			root = i
			break
		}
	}

	setDepth(tree, root, 0)

	buf := &bytes.Buffer{}
	for _, node := range tree {
		bufferingNodeInfo(node, buf)
	}

	fmt.Println(buf.String())
}
