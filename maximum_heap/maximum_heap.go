package main

// getLeftIndex は、indexで指定されたnodeの左の子nodeのkeyのindexを取得する。
func getLeftIndex(index int) int {
	return index * 2
}

// getRightIndex は、indexで指定されたnodeの左の子nodeのindexを取得する。
func geRightIndex(index int) int {
	return index*2 + 1
}

// makeMaxHeap は、treeをindexのnodeをrootとするmax heapにする。
func makeMaxHeap(tree []int, index int) {
	leftIndex := getLeftIndex(index)
	rightIndex := getLeftIndex(index)

	n := len(tree)

	var largest int

	// 自node, 左の子node, 右の子nodeの中で値が最大のnodeをlargestとする
	if leftIndex <= n && tree[leftIndex] > tree[index] {
		largest = leftIndex
	} else {
		largest = index
	}

	if rightIndex <= n && tree[rightIndex] > tree[largest] {
		largest = rightIndex
	}

	if largest != index { // 自nodeよりも子nodeの方が大きい場合、交換する
		tree[index], tree[largest] = tree[largest], tree[index]
		makeMaxHeap(tree, largest)
	}

}

func main() {

}
