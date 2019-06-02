package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// merge は、マージを行う。
func merge(list []int, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid

	lLen, rLen := n1+1, n2+1
	leftList, rightList := make([]int, lLen, lLen), make([]int, rLen, rLen)

	// 左側のSliceを作成する
	for i := 0; i < n1; i++ {
		leftList[i] = list[left+i]
	}

	// 番兵としてありえない数字を置く
	leftList[n1] = math.MaxInt32

	// 右側のSliceを作成する
	for i := 0; i < n2; i++ {
		rightList[i] = list[mid+i]
	}

	// 番兵としてありえない数字を置く
	rightList[n2] = math.MaxInt32

	i, j := 0, 0

	// 左側のSliceと右側のSliceを比較しながら元のSliceに格納していく
	for k := left; k < right; k++ {
		if leftList[i] <= rightList[j] {
			list[k] = leftList[i]
			i = i + 1
		} else {
			list[k] = rightList[j]
			j = j + 1
		}
	}
}

// mergeSort は、メージソートを行う。
func mergeSort(list []int, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		// 左側の整列
		mergeSort(list, left, mid)
		// 右側の整列
		mergeSort(list, mid, right)
		merge(list, left, mid, right)
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

}
