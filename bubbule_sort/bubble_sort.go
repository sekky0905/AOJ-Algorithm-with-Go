package main

import "fmt"

func main() {
	numSlice, length := getTarget()
	sortedSlice, count := BubbleSort(numSlice, length)
	// オペランドの間にスペースを挿入する
	formatted := fmt.Sprint(sortedSlice)
	fmt.Println(formatted[1 : len(formatted)-1])
	fmt.Println(count)
}

// getTarget は、ターゲットとなるintのSliceとその長さ返す
func getTarget() (target []int, length int) {
	fmt.Scan(&length)

	target = make([]int, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target, length
}

// BubbleSort は、バブルソートを行い数列を昇順に整列する
// 整列した数列及び、要素の交換回数を返す
func BubbleSort(numSlice []int, length int) ([]int, int) {
	isThereRightNext := true
	count := 0
	for isThereRightNext {
		// 交換の前にfalseにしておくことで、交換が行われなかった場合には上のループが終了することになる
		isThereRightNext = false
		for i := length - 1; i > 0; i-- {
			if numSlice[i] < numSlice[i-1] {
				tmp := numSlice[i]
				numSlice[i] = numSlice[i-1]
				numSlice[i-1] = tmp
				count++
				isThereRightNext = true
			}
		}
	}
	return numSlice, count
}
