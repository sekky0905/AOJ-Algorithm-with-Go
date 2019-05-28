package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// card は、カードを表す。
type card struct {
	suit string
	num  int
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

func scanToString() string {
	sc.Scan()
	return sc.Text()
}

// quickSort は、クリックソートを行う。
func quickSort(cards []*card, p, r int) {
	q := 0
	if p < r {
		q = partition(cards, p, r)
		// 左
		quickSort(cards, p, q-1)
		// 右
		quickSort(cards, q+1, r)
	}
}

// partition は、パーティションを使用しSortする。
// パーティションの対象は、p〜rとする。
func partition(cards []*card, p, r int) int {
	x := cards[r].num
	i := p - 1

	for j := p; j < r; j++ {
		if cards[j].num <= x {
			i++
			cards[i], cards[j] = cards[j], cards[i]
		}
	}
	cards[i+1], cards[r] = cards[r], cards[i+1]
	return i + 1
}

// merge は、マージを行う。
func merge(cards []*card, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid

	lLen, rLen := n1+1, n2+1
	leftCards, rightCards := make([]*card, lLen, lLen), make([]*card, rLen, rLen)

	// 左側のSliceを作成する
	for i := 0; i < n1; i++ {
		leftCards[i] = cards[left+i]
	}

	// 番兵としてありえない数字を置く
	dummy := &card{
		suit: "",
		num:  math.MaxInt32,
	}
	leftCards[n1] = dummy

	// 右側のSliceを作成する
	for i := 0; i < n2; i++ {
		rightCards[i] = cards[mid+i]
	}

	// 番兵としてありえない数字を置く
	rightCards[n2] = dummy

	i, j := 0, 0

	// 左側のSliceと右側のSliceを比較しながら元のSliceに格納していく
	for k := left; k < right; k++ {
		if leftCards[i].num <= rightCards[j].num {
			cards[k] = leftCards[i]
			i = i + 1
		} else {
			cards[k] = rightCards[j]
			j = j + 1
		}
	}
}

// mergeSort は、メージソートを行う。
func mergeSort(cards []*card, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		// 左側の整列
		mergeSort(cards, left, mid)
		// 右側の整列
		mergeSort(cards, mid, right)
		merge(cards, left, mid, right)
	}
}

func main() {
	n := scanToInt()

	a := make([]*card, n, n)
	b := make([]*card, n, n)
	for i := 0; i < n; i++ {
		cardStr := scanToString()
		s := strings.Split(cardStr, " ")

		suit := s[0]
		num, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}

		card1 := &card{
			suit: suit,
			num:  num,
		}

		card2 := card1

		a[i] = card1
		b[i] = card2
	}

	mergeSort(a, 0, n)
	quickSort(b, 0, n-1)

	if reflect.DeepEqual(a, b) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}

	for _, v := range b {
		fmt.Printf("%s %d\n", v.suit, v.num)
	}
}
