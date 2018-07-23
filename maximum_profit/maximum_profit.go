package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	MaximumProfit()
}

// getTarget は、ターゲットとなるintを返す
func getTarget(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

// getMax は、最大値を返す
func getMax(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// getMin は、最小値を返す
func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// MaximumProfit は、Rj-Ri(j > i)の中で最大利益を出力する
func MaximumProfit() {
	sc := bufio.NewScanner(os.Stdin)

	length := getTarget(sc)
	// とりあえず、先頭の数字をminVとおく
	minV := getTarget(sc)

	maxV := math.MinInt64
	for j := 1; j < length; j++ {
		x := getTarget(sc)
		// その時点でのMaxと(今回の数字-最小値)の大きい方がMax
		maxV = getMax(maxV, x-minV)
		// その時点のMinと今回の数字の小さい方がMin
		minV = getMin(x, minV)
	}

	fmt.Println(maxV)
}
