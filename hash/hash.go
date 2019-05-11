package main

func main() {

}

// hashTable は、ハッシュテーブルを表す。
type hashTable []string

// hash1 は、引数で与えられたkeyから、key mod len(hashTable)を計算する。
func (h *hashTable) hash1(key int) int {
	// len(hashTable) = mとする時に、0~m-1になるようにする。
	return key % len(*h)
}

// hash2 は、与えられたkeyとiから、1 + (key mod len(hashTable) -1) を計算する。
func (h *hashTable) hash2(key int) int {
	return 1 + (key % (len(*h) - 1))
}

// hash は、hashTableに対するハッシュ値を取得する。
func (h *hashTable) hash(key, i int) int {
	return (h.hash1(key) + i*h.hash2(key)) % len(*h)
}
