package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// n個の数字が並ぶ数列s
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	// q個の数字が並ぶ数列t(qは使用しないので省略)
	sc.Scan()
	_ = sc.Text()

	sc.Scan()
	t := strings.Split(sc.Text(), " ")

	count := 0

	for _, elm := range t {
		if LinearSearch(elm, s, n) {
			count++
		}
	}
	fmt.Println(count)
}

// LinearSearch は、線形探索を行う。
func LinearSearch(key string, target []string, length int) bool {
	i := 0

	// 番兵を置く
	target = append(target, key)

	// 番兵を置くので、ifを毎回行わなくても良い
	for target[i] != key {
		i++
		if i == length {
			return false
		}
	}
	return true
}
