package main

import "math"

// getParentIndex は、indexで指定されたnodeの親nodeのindexを取得する。
func getParentIndex(index int) int {
	return index / 2
}

// getLeftIndex は、indexで指定されたnodeの左の子nodeのkeyのindexを取得する。
func getLeftIndex(index int) int {
	return index * 2
}

// getRightIndex は、indexで指定されたnodeの左の子nodeのindexを取得する。
func getRightIndex(index int) int {
	return index*2 + 1
}

// isValidIndex は与えられたindexが有効なものかどうかを確認する。
func isValidIndex(index, n int) bool {
	return index <= n && index > 0
}

// makeMaxHeap は、treeをindexのnodeをrootとするmax heapにする。
func makeMaxHeap(tree []int, index int) {
	leftIndex := getLeftIndex(index)
	rightIndex := getRightIndex(index)

	n := len(tree)

	var largest int

	// 自node, 左の子node, 右の子nodeの中で値が最大のnodeをlargestとする
	if leftIndex < n && tree[leftIndex] > tree[index] {
		largest = leftIndex
	} else {
		largest = index
	}

	if rightIndex < n && tree[rightIndex] > tree[largest] {
		largest = rightIndex
	}

	if largest != index { // 自nodeよりも子nodeの方が大きい場合、交換する
		tmp := tree[index]
		tree[index] = tree[largest]
		tree[largest] = tmp
		makeMaxHeap(tree, largest)
	}

}

func buildMaxHeap(tree []int) {
	// 子nodeを持つ最大のnodeのindexは、len(tree) / 2
	for i := len(tree) / 2; i >= 1; i-- {
		makeMaxHeap(tree, i)
	}

}

// insert は、treeの適切な位置にkeyを格納する。
func insert(tree []int, key int) {
	tree = append(tree, math.MinInt32)
	heapIncreaseKey(tree, len(tree), key)
}

// heapIncreaseKey は、treeのindexで指定されたnode以上の適切な位置にkeyを格納する。
func heapIncreaseKey(tree []int, index, key int) {
	// keyが現在のkeyよりも小さい場合は、return
	if key < tree[index] {
		return
	}

	// keyをtreeに格納する
	tree[index] = key

	// 今回のnodeから、rootまで親を辿って遡る
	// 親よりも今回のkeyの方が大きい場合、親と入れ替える
	for index > 1 && tree[getParentIndex(index)] < tree[index] {
		tree[index], tree[getParentIndex(index)] = tree[getParentIndex(index)], tree[index]
		index = getParentIndex(index)
	}
}

func main() {

}
