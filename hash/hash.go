package main

import (
	"bufio"
	"bytes"
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

var length = 1000

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

// insert は、ハッシュテーブルにデータを格納する。
func insert(hashTable []string, str string) {
	i := 0
	key := getKeyFromString(str)
	for {
		hashed := hash(key, i)
		if hashTable[hashed] == "" { // 指定した箇所に何も入っていなかったら、格納
			hashTable[hashed] = str
			return
		}
		i++ // hashのiをインクリメントする
	}
}

// isExist は、引数で与えられた文字列がhashTable上に既に存在するかどうかを確認する。
func isExist(hashTable []string, str string) bool {
	i := 0
	key := getKeyFromString(str)
	for {
		hashed := hash(key, i)
		if hashTable[hashed] == str {
			return true
		} else if hashTable[hashed] == "" || i >= length { // i毎にhashでズラして進めていく中でから文字が存在したら、それ以上進めても存在しないので
			return false
		}
		i++
	}
}

// hash は、hashTableに対するハッシュ値を取得する。
func hash(key, i int) int {
	return (hash1(key) + i*hash2(key)) % length
}

// hash1 は、引数で与えられたkeyから、key mod len(hashTable)を計算する。
func hash1(key int) int {
	// len(hashTable) = mとする時に、0~m-1になるようにする。
	return key % length
}

// hash2 は、与えられたkeyとiから、1 + (key mod len(hashTable) -1) を計算する。
func hash2(key int) int {
	return 1 + (key % (length - 1))
}

// 1文字を数字に変換する。
func getIntFromChar(char byte) int {
	switch char {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	default:
		panic("invalid char")
	}
}

// getKeyFromString は文字列からkeyを得る。
func getKeyFromString(str string) int {
	const adjusterValue = 5
	sum, adjuster := 0, 1

	for i := range str {
		sum += adjuster * (getIntFromChar(str[i]))
		// 例えば、AATの場合6になる。
		// CCCの場合も6になり、衝突してしまうので各値に下駄を履かせる
		adjuster *= adjusterValue
	}
	return sum
}

func main() {
	length = scanToInt()
	// ハッシュテーブル
	hashTable := make([]string, length)

	buf := bytes.Buffer{}
	for i := 0; i < length; i++ {
		s := strings.Split(scanToString(), " ")
		if s[method] == "insert" {
			insert(hashTable, s[value])
		} else {
			if isExist(hashTable, s[value]) {
				buf.WriteString("yes\n")
			} else {
				buf.WriteString("no\n")
			}
		}
	}
	fmt.Println(strings.TrimRight(buf.String(), "\n"))
}
