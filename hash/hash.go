package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

const (
	method = iota
	value
)

func main() {
	n := scanToInt()
	h := hashTable{}

	for i := 0; i < n; i++ {
		str := scanToString()
		s := strings.Split(str, " ")
		if s[method] == "insert" {
			h.insert(s[value])
		} else {
			fmt.Print(boolToString(h.isExist(s[value])))
		}
	}
}

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func scanToString() string {
	sc.Scan()
	return sc.Text()
}

func boolToString(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// hashTable は、ハッシュテーブルを表す。
type hashTable map[int]string

// insert は、ハッシュテーブルにデータを格納する。
func (h hashTable) insert(str string) {
	i := 0
	key := getKeyFromString(str)
	for {
		hashed := h.hash(key, i)
		if h[hashed] == "" { // 指定した箇所に何も入っていなかったら、格納
			h[hashed] = str
			continue
		}
		i++ // hashのiをインクリメントする
	}
}

// isExist は、引数で与えられた文字列がhashTable上に既に存在するかどうかを確認する。
func (h hashTable) isExist(str string) bool {
	i := 0
	key := getKeyFromString(str)
	for {
		hashed := h.hash(key, i)
		if h[hashed] == str {
			return true
		} else if h[hashed] == "" || i >= len(h) { // i毎にhashでズラして進めていく中でから文字が存在したら、それ以上進めても存在しないので
			return false
		}
		i++
	}
}

// hash1 は、引数で与えられたkeyから、key mod len(hashTable)を計算する。
func (h hashTable) hash1(key int) int {
	// len(hashTable) = mとする時に、0~m-1になるようにする。
	return key % len(h)
}

// hash2 は、与えられたkeyとiから、1 + (key mod len(hashTable) -1) を計算する。
func (h hashTable) hash2(key int) int {
	return 1 + (key % (len(h) - 1))
}

// hash は、hashTableに対するハッシュ値を取得する。
func (h hashTable) hash(key, i int) int {
	return (h.hash1(key) + i*h.hash2(key)) % len(h)
}

// 1文字を数字に変換する。
func getIntFromChar(char string) int {
	switch char {
	case "A":
		return 1
	case "C":
		return 2
	case "G":
		return 3
	case "T":
		return 4
	default:
		panic("invalid char")
	}
}

// getKeyFromString は文字列からkeyを得る。
func getKeyFromString(str string) int {
	const adjusterValue = 5
	sum, adjuster := 0, 1

	for _, s := range str {
		sum += adjuster * (getIntFromChar(string(s)))
		// 例えば、AATの場合6になる。
		// CCCの場合も6になり、衝突してしまうので各値に下駄を履かせる
		adjuster *= adjusterValue
	}
	return sum
}
