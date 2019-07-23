package main

// sumAll は、与えられた数列を全て合計した値を返す。
func sumAll(target []int) int {
	result := 0
	for _, v := range target {
		result += v
	}
	return result
}

func main() {

}
