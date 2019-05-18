package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	n := scanToInt()
	k := scanToInt()

	w := make([]int, n, n)
	for i := 0; i < n; i++ {
		w[i] = scanToInt()
	}

	p := BinarySearch(n, k, w)
	fmt.Println(p)
}

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

// getMaxCargoNum は、与えられた積載量pでsliceで与えられた荷物(w)をk台以内のトラックで
// 積み込むことができる最大の個数を返す。
func getMaxCargoNum(n, k, p int, w []int) int {
	// k台のトラック全体の荷物の数合計をカウントするカウンター
	counterOfCargo := 0
	for i := 0; i < k; i++ { // トラックが0台~k台まで
		// トラック内での積み込んだ荷物の重さ合計
		sumOfCargoInTruck := 0
		for sumOfCargoInTruck+w[counterOfCargo] <= p { // これまでトラックに積んだ荷物の合計と今回の荷物を足したものが、与えられた積載量以下であれば処理を続ける
			sumOfCargoInTruck += w[counterOfCargo] // // トラック内での積み込んだ荷物の合計に今回の荷物を追加
			counterOfCargo++                       // 荷物の個数をカウントアップする
			if counterOfCargo == n {               //
				return n // k台のトラック全体の荷物の数合計と与えられた荷物の個数が同じということは、処理が終了
			}
		}
	}
	return counterOfCargo
}

const (
	maxN = 100000
	maxW = 10000
)

// BinarySearch は、2分探索を行う。
func BinarySearch(n, k int, w []int) int {
	left, right := 0, maxN*maxW

	for right > left {
		p := (left + right) / 2
		v := getMaxCargoNum(n, k, p, w)

		if v >= n {
			right = p // vがn以上になる最初のpを返す
		} else { // 探しているkeyが真ん中の値よりも大きい婆には、左端をp+1にし探索の幅を狭める
			left = p + 1
		}
	}

	return right
}
