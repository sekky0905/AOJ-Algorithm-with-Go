package main

import (
	"bufio"
	"bytes"
	"fmt"
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
		counter++ // 比較回数をカウントアップ
		if leftList[i] <= rightList[j] {
			list[k] = leftList[i]
			i = i + 1
		} else {
			list[k] = rightList[j]
			j = j + 1
		}
	}
}

var counter int

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
var buf bytes.Buffer

func main() {
	sc.Split(bufio.ScanWords)

	n := scanToInt()
	list := make([]int, n, n)
	for i := 0; i < n; i++ {
		list[i] = scanToInt()
	}

	mergeSort(list, 0, n)

	for i, v := range list {
		if i == n-1 {
			buf.WriteString(fmt.Sprintf("%d", v))
		} else {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
	}

	fmt.Println(buf.String())
	fmt.Println(counter)
}

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
