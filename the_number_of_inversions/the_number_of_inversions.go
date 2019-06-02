package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// merge は、マージを行う。
func merge(list []int, left, mid, right int) int {
	n1 := mid - left
	n2 := right - mid

	// 反転数のカウントを保存
	counter := 0

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
			counter += n1 - i
		}
	}
	return counter
}

// mergeSort は、メージソートを行う。
func mergeSort(list []int, left, right int) int {
	counter := 0
	if left+1 < right {
		mid := (left + right) / 2
		// 左側の整列
		counter += mergeSort(list, left, mid)
		// 右側の整列
		counter += mergeSort(list, mid, right)
		counter += merge(list, left, mid, right)
	}
	return counter
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

	list := make([]int, n, n)
	for i := 0; i < n; i++ {
		list[i] = scanToInt()
	}

	counter := mergeSort(list, 0, n)

	fmt.Println(counter)
}
