package main

import (
	"bufio"
	"os"
	"strconv"
)

type orders []int

// preOrders は、先行順巡回による数列
var preOrders orders

// inOrders は、中間順巡回による数列
var inOrders orders

// find は、引数で与えられた数字のordersでのindexを返す。
func (o orders) find(target int) int {
	for i, v := range o {
		if v == target {
			return i
		}
	}
	return -1
}

func getInputList(n int) []int {
	list := make([]int, n, n)
	for j := 0; j < n; j++ {
		list[j] = scanToInt()
	}
	return list
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
	preOrders = getInputList(n)
	inOrders = getInputList(n)

}
