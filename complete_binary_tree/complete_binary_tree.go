package main

type tree []int

// getParentIndex は、indexで指定されたnodeの親nodeのindexを取得する。
func getParent(index int) int {
	return index / 2
}

// getLeft は、indexで指定されたnodeの左の子nodeのkeyのindexを取得する。
func getLeftIndex(index int) int {
	return index * 2
}

// getRight は、indexで指定されたnodeの左の子nodeのindexを取得する。
func geRight(index int) int {
	return index*2 + 1
}

func main() {

}
