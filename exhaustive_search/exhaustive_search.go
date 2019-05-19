package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type targetSlice []int

// canCreateTargetInt は、ターゲットとなる整数をSliceの数字を使って作れるかどうかを確認する。
func (t targetSlice) canCreateTargetInt(i, m int) bool {
	if m == 0 {
		return true
	}

	if i >= len(t) {
		return false
	}

	// t[i]を使わない場合(左)と使う場合(右)に分けてシミュレーションする。
	result := t.canCreateTargetInt(i+1, m) || t.canCreateTargetInt(i+1, m-t[i])
	return result
}

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func printBool(b bool) {
	if b {
		fmt.Println("yes")
		return
	}
	fmt.Println("no")
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanToInt()

	A := make(targetSlice, n, n)
	for i := 0; i < n; i++ {
		A[i] = scanToInt()
	}

	q := scanToInt()
	for i := 0; i < q; i++ {
		m := scanToInt()
		result := A.canCreateTargetInt(0, m)
		printBool(result)
	}
}
