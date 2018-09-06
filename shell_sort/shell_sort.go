package main

import (
	"fmt"
	"strings"
)

func main() {
	target, length := getTarget()
	sorted, g, cnt := ShellSort(target, length)
	fmt.Println(len(g))
	printHorizontal(g)
	fmt.Println(cnt)
	printVertical(sorted)
}

// getTarget は、ターゲットとなるintのSliceとその長さ返す
func getTarget() (target []int, length int) {
	fmt.Scan(&length)

	target = make([]int, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target, length
}

// InsertionSort は、挿入ソートを実現する
func InsertionSort(target []int, length, g int) ([]int, int) {
	cnt := 0
	// gから、最後までを整列させる
	for i := g; i < length; i++ {
		tmp := target[i] // これを軸とする
		// iとgだけ離れているindexをjとする
		j := i - g
		for j >= 0 && target[j] > tmp {
			// j+gにjを入れる
			// j+gは要するにi
			target[j+g] = target[j]
			// jを減らしていくので、↑はずれる
			j = j - g
			cnt++
		}
		target[j+g] = tmp // 今回の軸が入るのにふさわしいところに入れる
	}
	return target, cnt
}

// ShellSort はShell Sortを実現する
func ShellSort(target []int, length int) ([]int, []int, int) {
	var g []int

	for h := 1; h <= length; h = 3*h + 1 {
		g = append(g, h)
	}

	cnt := 0
	m := 0
	for i := len(g) - 1; 0 <= i; i-- {
		target, m = InsertionSort(target, length, g[i])
		cnt = cnt + m
	}

	return target, g, cnt
}

// printVertical は、垂直に数字を標準出力する
func printVertical(target []int) {
	trimmed := strings.Trim(fmt.Sprint(target), "[]")
	fmt.Println(strings.Replace(trimmed, " ", "\n", -1))
}

// printHorizontal は、水平に数字を標準出力する
func printHorizontal(target []int) {
	fmt.Println(strings.Trim(fmt.Sprint(target), "[]"))
}
