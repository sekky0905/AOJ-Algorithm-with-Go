package main

// node は、木の接点もしくは葉を表す。
type node struct {
	id       int
	parent   int
	depth    int
	nodeType string
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

func main() {

}
