package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type orders []int

// find は、引数で与えられた数字のordersでのindexを返す。
func (o orders) find(target int) int {
	for i, v := range o {
		if v == target {
			return i
		}
	}
	return -1
}

// print は、ordersの要素を出力する。
func (o orders) print() {
	for i, v := range o {
		if i == len(o)-1 {
			fmt.Println(v)
			return
		}
		fmt.Printf("%d ", v)
	}
}

// preOrders は、先行順巡回による数列
var preOrders orders

// inOrders は、中間順巡回による数列
var inOrders orders

// originalOrders は、復元される数列
var originalOrders orders

var buf = bytes.Buffer{}

// preCursor は、preOrdersの現在位置
var preCursor = -1

// reconstruct は、先行順巡回及び中間順巡回による数列から、二分木を再構築する。
func reconstruct(left, right int) {
	if left >= right {
		return
	}

	// 先行順巡回による数列の現在位置を1つ進める
	preCursor++

	// 先行順巡回による数列の次の節点を取得する
	c := preOrders[preCursor]
	// 中間順巡回による数列の中での次の接点の位置を取得する
	// 節点は、左部分木と右部分木の間の中間地点になる
	m := inOrders.find(c)

	// 左部分木の再構築を行う
	// 左端から節点まで
	reconstruct(left, m)
	// 右部分木の再構築を行う
	// 節点から右端まで
	reconstruct(m+1, right)

	originalOrders = append(originalOrders, c)
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

	left, right := 0, len(preOrders)
	reconstruct(left, right)
	originalOrders.print()
}
