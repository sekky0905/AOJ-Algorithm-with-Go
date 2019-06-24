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

func main() {

}
