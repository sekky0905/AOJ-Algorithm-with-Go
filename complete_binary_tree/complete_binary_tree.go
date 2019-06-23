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

		// index=0は使用しないので、スキップ
		if i == 0 {
			continue
		}

		parent, left, right := getParent(i), getLeftIndex(i), geRightIndex(i)

		buf.WriteString(fmt.Sprintf("node %d: key = %d, ", i, node))

		if isValidIndex(parent, n) {
			buf.WriteString(fmt.Sprintf("parent key = %d, ", tree[parent]))
		}

		if isValidIndex(left, n) {
			buf.WriteString(fmt.Sprintf("left key = %d, ", tree[left]))
		}

		if isValidIndex(right, n) {
			buf.WriteString(fmt.Sprintf("right key = %d, ", tree[right]))
		}
		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
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
	sc.Split(bufio.ScanWords)
	n := scanToInt()
	// index=0は、使用しないため
	tree := make([]int, n+1, n+1)
	for i := 1; i < n+1; i++ {
		v := scanToInt()
		tree[i] = v
	}

	print(tree, n)
}
