package main

import (
	"fmt"
)

func main() {
	// targetを取得する
	target, length := getTarget()
	// ソートする
	insertionSort(target, length)
}

// getTarget は、ターゲットとなるintのSliceを返す
func getTarget() (target []int, length int) {
	fmt.Scan(&length)

	target = make([]int, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target, length
}

// insertionSort は、挿入ソートを行なう
func insertionSort(target []int, length int) {
	for i := 1; i < length; i++ {
		tmp := target[i] // これを軸とする
		j := i - 1
		for j >= 0 && target[j] > tmp {
			target[j+1] = target[j]
			j--
		}
		target[j+1] = tmp // 今回の軸が入るのにふさわしいところに入れる
		print(target)
	}
}

// print は、数字を標準出力する
func print(target []int) {
	length := len(target)
	for i, v := range target {
		if i == length-1 {
			fmt.Printf("%d\n", v)
		} else {
			fmt.Printf("%d", v)
		}
	}
}
