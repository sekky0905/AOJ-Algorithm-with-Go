package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// getParentIndex は、indexで指定されたnodeの親nodeのindexを取得する。
func getParent(index int) int {
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

func print(tree []int, n int) {
	var buf bytes.Buffer

	for i, node := range tree {
		p, r, l := getParent(i), getLeftIndex(i), geRightIndex(i)

		buf.WriteString(fmt.Sprintf("node %d: key = %d,", i, node))
		if isValidIndex(p, n) {
			buf.WriteString(fmt.Sprintf(" parent key =%d,", p))
		}

		if isValidIndex(r, n) {
			buf.WriteString(fmt.Sprintf(" left key =%d,", p))
		}

		if isValidIndex(l, n) {
			buf.WriteString(fmt.Sprintf(" right key = =%d", p))
		}
		buf.WriteString("\n")
	}
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

func main() {
	n := scanToInt()

	sc.Split(bufio.ScanWords)

	// index=0は、使用しないため
	tree := make([]int, n+1, n+1)
	for i := 0; i < n; i++ {
		v := scanToInt()
		tree[i] = v
	}

}
