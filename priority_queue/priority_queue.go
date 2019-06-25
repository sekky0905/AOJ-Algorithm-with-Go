package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

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

const (
	rootIndex   = 1
	methodIndex = 0
	keyIndex    = 1
)

// getMaxFromHeap は、heapの最大要素を取得する。
func getMaxFromHeap(tree []int) int {
	return tree[rootIndex]
}

func deleteMaxFromHeap(tree []int) error {
	if len(tree) < 1 {
		return errors.New("tree should be over 0")
	}

	n := len(tree)
	tree[rootIndex] = tree[n]
	tree = delete(tree, n)
	makeMaxHeap(tree, rootIndex)

	return nil
}

func delete(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

var sc = bufio.NewScanner(os.Stdin)

func scanToText() string {
	sc.Scan()
	return sc.Text()
}

func solve(tree []int, method, keyStr string) error {
	key, err := strconv.Atoi(keyStr)
	if err != nil {
		return err
	}

	if method == "insert" {
		insert(tree, key)
	}

	if method == "extract" {
		max := getMaxFromHeap(tree)
		if err := deleteMaxFromHeap(tree); err != nil {
			return err
		}
		fmt.Println(max)
	}
	return nil
}

func main() {
	sc.Split(bufio.ScanWords)

	tree := make([]int, 0)

	for {
		str := scanToText()
		if str == "end" {
			break
		}
		s := strings.Split(str, " ")
		m, k := s[methodIndex], s[keyIndex]
		if err := solve(tree, m, k); err != nil {
			panic(err)
		}
	}
}
