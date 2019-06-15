package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	n := scanToInt()

	// 先行順巡回による数列
	pre := getInputList(n)
	// 中間順巡回による数列
	in := getInputList(n)

}

func getInputList(n int) []int {
	list := make([]int, n, n)
	for j := 0; j < n; j++ {
		list[j] = scanToInt()
	}
	return list
}

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
