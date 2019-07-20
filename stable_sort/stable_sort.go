package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// card は、カードを表す
type card struct {
	rowVal   string
	alphabet string
	num      int
}

// getTarget は、ターゲットとなるintのSliceとその長さ返す
func getTarget() (target []string, length int) {
	fmt.Scan(&length)

	target = make([]string, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target, length
}

// newCard は、cardを生成し、返す
func newCard(target string) *card {
	num, _ := strconv.Atoi(target[1:])
	return &card{
		rowVal:   target,
		alphabet: target[:1],
		num:      num,
	}
}

// BubbleSort は、バブルソートを行い数列を昇順に整列する
// 整列した数列及び、要素の交換回数を返す
func BubbleSort(target []*card, length int) []*card {
	isThereRightNext := true
	for isThereRightNext {
		// 交換の前にfalseにしておくことで、交換が行われなかった場合には上のループが終了することになる
		isThereRightNext = false
		for i := length - 1; i > 0; i-- {
			if target[i].num < target[i-1].num {
				tmp := target[i]
				target[i] = target[i-1]
				target[i-1] = tmp
				isThereRightNext = true
			}
		}
	}
	return target
}

// SelectionSort は選択ソートを実現する
func SelectionSort(target []*card, length int) []*card {
	// i は、未ソート部分の先頭
	for i := 0; i < length; i++ {
		// とりあえず、今回の数字をminとおく
		min := target[i]
		minIndex := i

		// j は、未ソートをWalkする
		for j := i; j < length; j++ {
			// jの値が、minより小さい場合は、minにjの値を格納する
			if target[j].num < min.num {
				min = target[j]
				minIndex = j
			}
		}

		// 未ソート部分の先頭の場所と最小の場所が同じ場合には、交換しない
		if minIndex != i {
			target[i], target[minIndex] = target[minIndex], target[i]
		}
	}

	return target
}

// isStableCalculatedByFastAlgorithm は、高速で、Sortしたアルゴリズムが、安定的かどうかを確認する
// Bubble Sortが安定的であると言う性質を利用して、Bubble SortでSort済みのSliceと同じであれば安定的とみなす
func isStableCalculatedByFastAlgorithm(afterSorted, bubbleSorted []*card) bool {
	return reflect.DeepEqual(afterSorted, bubbleSorted)
}

// isStableCalculatedBySlowAlgorithm は、入力と出力からそのSliceをSortしたアルゴリズムが、安定的かどうかを確認する
// beforeSortedは、Sort前のSliceを表し、afterSortedは、Sort後のSliceを表す
// 遅いので、別の方法の方がいい
func isStableCalculatedBySlowAlgorithm(beforeSorted, afterSorted []*card) bool {
	length := len(beforeSorted)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for a := 0; a < length; a++ {
				for b := a + 1; b < length; b++ {
					// Sort前の隣同士の数字が同じでかつ、前後で文字まで合わせた時に順番が変わってたら、NotStable
					if beforeSorted[i].num == beforeSorted[j].num && beforeSorted[i].rowVal == afterSorted[b].rowVal && beforeSorted[j].rowVal == afterSorted[a].rowVal {
						return false
					}
				}
			}
		}
	}
	return true
}

// printIsStable は、Stableかどうかを標準出力する
func printIsStable(isStable bool) {
	if isStable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

// printSlice は、数字を標準出力する
func printSlice(target []*card, length int) {
	rows := make([]string, length, length)
	for i, v := range target {
		rows[i] = v.rowVal
	}

	formatted := fmt.Sprint(rows)
	fmt.Println(formatted[1 : len(formatted)-1])
}

func main() {
	// targetを取得する
	target, length := getTarget()

	targetCards := make([]*card, length, length)

	for i, v := range target {
		c := newCard(v)
		targetCards[i] = c
	}

	// 単純に代入するだけだと参照渡しになってしまうため、copy
	targetCards2 := make([]*card, length, length)
	copy(targetCards2, targetCards)

	// BubbleSort
	bSorted := BubbleSort(targetCards, length)
	printSlice(bSorted, length)
	//printIsStable(isStableCalculatedBySlowAlgorithm(targetCards, bSorted))
	printIsStable(isStableCalculatedByFastAlgorithm(targetCards, bSorted))

	// SelectionSort
	sSorted := SelectionSort(targetCards2, length)
	printSlice(sSorted, length)
	//printIsStable(isStableCalculatedBySlowAlgorithm(targetCards, sSorted))
	printIsStable(isStableCalculatedByFastAlgorithm(targetCards2, bSorted))
}
