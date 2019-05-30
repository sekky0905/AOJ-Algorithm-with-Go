package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func countingSort(a []int, n int) []int {
	// 0がn個存在するsliceを生成する
	c := make([]int, n, n)

	for j := 1; j < n; j++ {
		// aの数列のj番目に入ってる数字 = xとする
		// c[x]を++する
		// こうすることで、c[x]にxがaに存在する個数をカウントしていく
		c[a[j]]++
	}

	for i := 1; i < n; i++ {
		// c[i]にi以下の出現個数の累計数を記録する
		c[i] = c[i] + c[i-1]
	}

	// 整列後の数列の出力先
	b := make([]int, n, n)

	for j := n; j > 1; j-- {
		b[c[a[j]]] = a[j]
		c[a[j]]-- // a[j]番目のcountの累計個数を減らす
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

	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = scanToInt()
	}

	b := countingSort(a, n)

	buf := bytes.Buffer{}

	for i, v := range b {
		if i == len(b)-1 {
			buf.WriteString(fmt.Sprintf("%d\n", v))
		} else {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
	}

	fmt.Println(buf.String())
}
