package main

import (
	"bufio"
	"os"
	"strconv"
)

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

func buildMaxHeap(tree []int) {
	// 子nodeを持つ最大のnodeのindexは、len(tree) / 2
	for i := len(tree) / 2; i >= 1; i-- {
		makeMaxHeap(tree, i)
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
	sc.Split(bufio.ScanWords)
	n := scanToInt()
	// index=0は、使用しないため
	tree := make([]int, n+1, n+1)
	for i := 1; i < n+1; i++ {
		v := scanToInt()
		tree[i] = v
	}
}
