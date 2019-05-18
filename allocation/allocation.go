package main

// getMaxCargoNum は、与えられた積載量pでsliceで与えられた荷物(w)をk台以内のトラックで
// 積み込むことができる最大の個数を返す。
func getMaxCargoNum(n, k, p int, w []int) int {
	// k台のトラック全体の荷物の数合計をカウントするカウンター
	counterOfCargo := 0
	for j := 0; j < k; j++ { // トラックが0台~k台まで
		// トラック内での積み込んだ荷物の重さ合計
		sumOfCargoInTruck := 0
		cargo := w[counterOfCargo]
		for sumOfCargoInTruck+cargo <= p { // これまでトラックに積んだ荷物の合計と今回の荷物を足したものが、与えられた積載量以下であれば処理を続ける
			sumOfCargoInTruck += cargo // // トラック内での積み込んだ荷物の合計に今回の荷物を追加
			counterOfCargo++           // 荷物の個数をカウントアップする
			if counterOfCargo == n {   //
				return counterOfCargo // k台のトラック全体の荷物の数合計と与えられた荷物の個数が同じということは、処理が終了
			}
		}
	}
	return counterOfCargo
}

// getMax
