package main

import (
	"bufio"
	"os"
	"strconv"
)

func calcMatrixChainMultiplication(n int, p []int, m [][]int) [][]int {
	for i := 1; i <= n; i++ {
		m[i][i] = 0
	}

	for l := 2; l <= n; l++ {
		for i := 1; i <= n-1; i++ {
			j := i + l - 1
			m[i][j] = 1 << 21
			for k := i; k <= j-1; k++ {
				m[i][j] = min(m[i][j], m[i][k]+m[k+1][j]+p[i-1]*p[k]*p[j])
			}
		}
	}
	return m
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

}
