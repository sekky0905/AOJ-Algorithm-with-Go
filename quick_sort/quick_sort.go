package main

import (
	"bufio"
	"os"
	"strconv"
)

// card は、カードを表す。
type card struct {
	suit string
	Num  int
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
	x := cards[r].Num
	i := p - 1

	for j := p; j < r; j++ {
		if cards[j].Num <= x {
			i++
			cards[i], cards[j] = cards[j], cards[i]
		}
	}
	cards[i+1], cards[r] = cards[r], cards[i+1]
	return i + 1
}

func main() {
	n := scanToInt()

	list := make([]int, n, n)
	for i := 0; i < n; i++ {
		list[i] = scanToInt()
	}
}
