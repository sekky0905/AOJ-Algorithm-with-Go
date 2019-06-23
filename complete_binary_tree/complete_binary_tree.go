package main

type tree []int

// getParentIndex は、indexで指定されたnodeの親nodeのindexを取得する。
func getParentIndex(index int) int {
	return index / 2
}

// getLeft は、indexで指定されたnodeの左の子nodeのkeyを取得する。
func (t tree) getLeft(index int) int {
	return t[index*2]
}

// getRight は、indexで指定されたnodeの左の子nodeのkeyを取得する。
func geRight() int {
	return 0
}

func main() {

}
