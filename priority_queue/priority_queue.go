package main

// getParentIndex は、indexで指定されたnodeの親nodeのindexを取得する。
func getParentIndex(index int) int {
	return index / 2
}

// getLeftIndex は、indexで指定されたnodeの左の子nodeのkeyのindexを取得する。
func getLeftIndex(index int) int {
	return index * 2
}

// getRightIndex は、indexで指定されたnodeの左の子nodeのindexを取得する。
func geRightIndex(index int) int {
	return index*2 + 1
}

// isValidIndex は与えられたindexが有効なものかどうかを確認する。
func isValidIndex(index, n int) bool {
	return index <= n && index > 0
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
