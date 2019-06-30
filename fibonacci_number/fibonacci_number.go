package fibonacci_number

import (
	"bufio"
	"os"
	"strconv"
)

var tree []int

const initialNum = 0

// calcFibonacci は、fibonacci数列のn番目の値を返す。
func calcFibonacci(n int) int {
	if n == 0 || n == 1 { // fibonacciにおいて、1番目、2番目は1
		tree[n] = 1
		return tree[n]
	}

	if tree[n] != initialNum { // 計算済みならその値を返す
		return tree[n]
	}

	tree[n] = calcFibonacci(n-2) + calcFibonacci(n-1) // 2個前と1個前を足す
	return tree[n]
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
