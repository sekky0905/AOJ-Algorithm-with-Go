package main

type tree []int

// getParent は、indexで指定されたnodeの親nodeのkeyを取得する。
func (t tree) getParent(index int) int {
	return t[index/2]
}

// getLeft は、indexで指定されたnodeの左の子nodeのkeyを取得する。
func getLeft() int {
	return 0
}

// getRight は、indexで指定されたnodeの左の子nodeのkeyを取得する。
func geRight() int {
	return 0
}

func main() {

}
