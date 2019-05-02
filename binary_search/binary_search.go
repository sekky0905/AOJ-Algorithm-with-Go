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

	// n個の数字が並ぶ数列s (nは使用しないので省略)
	n := scanToInt()

	// 今回は整列済みのデータが入ってくることが前提なので、整列は行わない。
	s := make([]int, n, n)
	for i := 0; i < n; i++ {
		s[i] = scanToInt()
	}

	q := scanToInt()
	count := 0

	for j := 0; j < q; j++ {
		t := scanToInt()
		if BinarySearch(s, t) {
			count++
		}
	}

	fmt.Println(count)
}

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

// BinarySearch は、2分探索を行う。
func BinarySearch(list []int, key int) bool {
	left, right := 0, len(list)

	for left < right {
		mid := (left + right) / 2

		if list[mid] == key {
			return true
		} else if key < list[mid] { // 探しているkeyが真ん中の値よりも小さい婆には、右端をmidにし探索の幅を狭める
			right = mid
		} else { // 探しているkeyが真ん中の値よりも大きい婆には、左端をmid+1にし探索の幅を狭める
			left = mid + 1
		}
	}

	return false
}
