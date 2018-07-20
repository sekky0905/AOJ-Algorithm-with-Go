package main

import (
	"fmt"
	"math"
)

// 与えられたn行の数字の中で素数の個数を表示する
func main() {
	numSlice := getTarget()
	count := 0
	for _, num := range numSlice {
		if IsPrimeNumber(num) {
			count++
		}
	}
	fmt.Println(count)
}

// IsPrimeNumber は、引数で与えられた数字が素数かどうかの真偽値を返す
// p ≤ √x を満たす素因子pが存在する場合、それは合成数(素数ではない)と判定する
func IsPrimeNumber(x int) bool {
	// 2は素数なので、2だった時点で返す
	if x == 2 {
		return true
	}

	// 素数は自然のみ対象となる
	// 2以外
	if x < 2 || x%2 == 0 {
		return false
	}

	// 2以下は対象外だから
	i := 3.0
	// xの平方根より小さい間は、iを増やし続け、それで割る
	for i <= math.Sqrt(float64(x)) {
		if x%int(i) == 0 {
			return false
		}
		i += 2
	}

	return true
}

// getTarget は、ターゲットとなるintのSliceを返す
func getTarget() []int {
	var length int
	fmt.Scan(&length)

	target := make([]int, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target
}
