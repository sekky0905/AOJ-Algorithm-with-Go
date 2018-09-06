package main

import (
	"fmt"
	"sort"
)

func main() {
	numSlice := getTarget()
	gcf := EuclideanAlgorithm(numSlice[0], numSlice[1])
	fmt.Println(gcf)
}

// getTarget は、ターゲットとなるintのSliceを返す
func getTarget() (target []int) {
	length := 2
	target = make([]int, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target
}

// EuclideanAlgorithm はユーグリッドの互除法を実現する
// A mod B = C
// Cは、AをBで割った余り
// A mod B = Cにおいて、Aを大きい方とし、Bを小さい方とする
// この次に、B mod C = Dを行う
// さらに
// C mod D = E と行った形で、再帰的に行う
// 割り切れた時に、最後の割る数が最大公約数になる
func EuclideanAlgorithm(x, y int) int {
	numSlice := []int{x, y}
	// 小さい順にsort
	sort.Slice(numSlice, func(i, j int) bool { return numSlice[i] > numSlice[j] })
	// 大きい方 / 小さい方 = 余りを行う
	big, little := numSlice[0], numSlice[1]
	result := big % little
	// 余りが0になった時点で、小さい方を返す
	if result != 0 {
		// 小さい方 / 余りを再帰的に行う
		little = EuclideanAlgorithm(little, result)
	}
	return little
}
