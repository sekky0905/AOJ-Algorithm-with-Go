package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	// realElmNum は、実際のtreeの中の要素の個数
	// indexの計算のためにtreeを表すsliceは実際の個数よりも多くなっている
	realElmNum = 0
	tree       []int
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

// makeMaxHeap は、treeをindexのnodeをrootとするmax heapにする。
func makeMaxHeap(index int) {
	leftIndex := getLeftIndex(index)
	rightIndex := getRightIndex(index)

	var largest int

	// 自node, 左の子node, 右の子nodeの中で値が最大のnodeをlargestとする
	if leftIndex <= realElmNum && tree[leftIndex] > tree[index] {
		largest = leftIndex
	} else {
		largest = index
	}

	if rightIndex <= realElmNum && tree[rightIndex] > tree[largest] {
		largest = rightIndex
	}

	if largest != index { // 自nodeよりも子nodeの方が大きい場合、交換する
		tree[index], tree[largest] = tree[largest], tree[index]
		makeMaxHeap(largest)
	}
}

// insert は、treeの適切な位置にkeyを格納する。
func insert(key int) {
	realElmNum++
	tree[realElmNum] = math.MinInt32
	heapIncreaseKey(realElmNum, key)
}

// heapIncreaseKey は、treeのindexで指定されたnode以上の適切な位置にkeyを格納する。
func heapIncreaseKey(index, key int) {
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
	rootIndex = 1
	keyIndex  = 1
)

// getMaxFromHeap は、heapの最大要素を取得する。
func getMaxFromHeap(tree []int) int {
	return tree[rootIndex]
}

func deleteMaxFromHeap(tree []int) error {
	if realElmNum < 1 {
		return errors.New("tree should be over 0")
	}

	// rootが削除対象なので、末尾のものをrootの場所に入れ、max heapに並び替えることでrootを削除したことを表す
	tree[rootIndex] = tree[realElmNum]
	realElmNum--
	makeMaxHeap(rootIndex)

	return nil
}

var sc = bufio.NewScanner(os.Stdin)

func scanToText() string {
	sc.Scan()
	return sc.Text()
}

func solve(str string) error {
	if str == "extract" {
		max := getMaxFromHeap(tree)
		if err := deleteMaxFromHeap(tree); err != nil {
			return err
		}
		fmt.Println(max)
	} else {
		s := strings.Split(str, " ")
		keyStr := s[keyIndex]
		key, err := strconv.Atoi(keyStr)
		if err != nil {
			return err
		}
		insert(key)
	}

	return nil
}

func main() {
	const max = 2000000
	tree = make([]int, max+1, max+1)
	for {
		str := scanToText()
		if str == "end" {
			break
		}
		if err := solve(str); err != nil {
			panic(err)
		}
	}
}
