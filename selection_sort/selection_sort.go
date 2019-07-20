package main

import "fmt"

// getTarget は、ターゲットとなるintのSliceを返す
func getTarget() (target []int, length int) {
	fmt.Scan(&length)

	target = make([]int, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target, length
}

// SelectionSort は選択ソートを実現する
func SelectionSort(target []int, length int) ([]int, int) {
	counter := 0

	// i は、未ソート部分の先頭
	for i := 0; i < length; i++ {
		// とりあえず、今回の数字をminとおく
		min := target[i]
		minIndex := i

		// j は、未ソートをWalkする
		for j := i; j < length; j++ {
			// jの値が、minより小さい場合は、minにjの値を格納する
			if target[j] < min {
				min = target[j]
				minIndex = j
			}
		}

		// 未ソート部分の先頭の場所と最小の場所が同じ場合には、交換しない
		if minIndex != i {
			target[i], target[minIndex] = target[minIndex], target[i]
			counter++
		}
	}

	return target, counter
}

func main() {
	// targetを取得する
	target, length := getTarget()
	// ソートする
	sortedSlice, count := SelectionSort(target, length)

	formatted := fmt.Sprint(sortedSlice)
	fmt.Println(formatted[1 : len(formatted)-1])
	fmt.Println(count)
}
