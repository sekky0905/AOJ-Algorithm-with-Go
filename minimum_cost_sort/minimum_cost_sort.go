package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const max = 100000

func getSorted(n int, target []int) []int {
	s := make([]int, n, n)
	copy(s, target)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func getIndexReversed(target []int) []int {
	s := make([]int, max)
	for i, v := range target {
		s[v] = i
	}
	return s
}

func minimumCostSort(n, x int, target []int) int {
	res := 0
	// 数列を正しく昇順に並び替えたもの
	sorted := getSorted(n, target)
	// indexと実態が逆の数列
	// 例)
	// 通常 {4, 1, 2, 5, 3}
	// これ {x, 1, 2, 4, 0, 3}
	indexReverse := getIndexReversed(sorted)
	completed := make([]bool, n, n)

	for i := range target {
		if completed[i] {
			continue
		}

		// サイクル内でのカーソル、合計、最小値、答え
		cursor, sum, minV, ans := i, 0, max, 0

		for {
			completed[cursor] = true
			ans++
			v := target[cursor]
			minV = min(minV, v)
			// 合計を今回の値分増やす
			sum += v
			cursor = indexReverse[v] // 値からindexを逆引き
			if completed[cursor] {
				break
			}
		}
		res1 := calc1(sum, ans, minV)
		res2 := calc2(sum, ans, minV, x)
		res += min(res1, res2)
	}
	return res
}

func calc1(sum, ans, minV int) int {
	return sum + (ans-2)*minV
}

func calc2(sum, ans, minV, x int) int {
	return sum + minV + (ans+1)*x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func main() {
	sc.Split(bufio.ScanWords)

	n := scanToInt()
	target := make([]int, n, n)
	x := max

	for i := 0; i < n; i++ {
		target[i] = scanToInt()
		x = min(x, target[i])
	}
	res := minimumCostSort(n, x, target)
	fmt.Println(res)
}
