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
func setDepth(nodes []node, currentNodeID, depth int) {
	nodes[currentNodeID].depth = depth
	for i := range nodes[currentNodeID].children {
		// 自分の child に当たる node の depth をそれぞれ再帰的に設定していく
		setDepth(nodes, nodes[currentNodeID].children[i], depth+1)
	}
}

func bufferingNodeInfo(node node, buf *bytes.Buffer) {
	nodeType := "leaf"
	if node.parent == -1 {
		nodeType = "root"
	} else if len(node.children) > 0 {
		nodeType = "internal node"
	}

	childStr := strings.Replace(fmt.Sprintf("%v", node.children), " ", ", ", -1)
	buf.WriteString(fmt.Sprintf("node %d: parent = %d, depth = %d, %s, %s\n", node.id, node.parent, node.depth, nodeType, childStr))
}

func initTree(nodes []node) {
	for i := range nodes {
		nodes[i].parent = -1
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scanToStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	// 高さ
	var height int
	fmt.Scan(&height)
	// 高さ height の tree
	tree := make([]node, height)

	initTree(tree)

	sc.Split(bufio.ScanWords)
	for i := 0; i < height; i++ {
		id, err := strconv.Atoi(scanToStr())
		if err != nil {
			panic(err)
		}

		degree, err := strconv.Atoi(scanToStr())
		if err != nil {
			panic(err)
		}

		// child node を埋める
		for j := 0; j < degree; j++ {
			c, err := strconv.Atoi(scanToStr())
			if err != nil {
				panic(err)
			}

			tree[id].children = append(tree[id].children, c)
			tree[c].parent = id
		}
	}

	root := -1
	for i := 0; i < height; i++ {
		if tree[i].parent == -1 {
			root = i
			break
		}
	}

	setDepth(tree, root, 0)

	buf := &bytes.Buffer{}
	for i, node := range tree {
		node.id = i
		bufferingNodeInfo(node, buf)
	}

	fmt.Print(buf.String())
}
