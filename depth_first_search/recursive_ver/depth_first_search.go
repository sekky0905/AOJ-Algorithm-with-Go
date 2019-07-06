package recursive_ver

const maxLength = 100

var (
	// adjacentMatrix は、隣接行列を表す。
	adjacentMatrix [][]int
	// timeCounter は全体の時刻を表す。
	timeCounter int
)

type color string

const (
	white color = "WHITE"
	gray  color = "GRAY"
	black color = "BLACK"
)

// node は、グラフにおける頂点を表す。
type node struct {
	color
	foundTime     int
	completedTime int
}

func main() {

}
