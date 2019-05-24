package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

type node struct {
	x float64
	y float64
}

// 60度をラジアンに変換する
const theta = math.Pi * 60.0 / 180.0

// calcKochCurve は、コッホ曲線の計算を行う。
func calcKochCurve(n int, p1, p2 *node) {
	if n == 0 { // 再帰の回数の残高がなくなったら終わり
		return
	}

	// _|__の位置
	s := &node{
		x: (2.0*p1.x + 1.0*p2.x) / 3.0,
		y: (2.0*p1.y + 1.0*p2.y) / 3.0,
	}

	// __|_の位置
	t := &node{
		x: (1.0*p1.x + 2.0*p2.x) / 3.0,
		y: (1.0*p1.y + 2.0*p2.y) / 3.0,
	}

	// 頂点
	u := &node{
		x: s.x + (t.x-s.x)*math.Cos(theta) - (t.y-s.y)*math.Sin(theta),
		y: s.y + (t.x-s.x)*math.Sin(theta) + (t.y-s.y)*math.Cos(theta),
	}

	calcKochCurve(n-1, p1, s)
	buffering(s)
	calcKochCurve(n-1, s, u)
	buffering(u)
	calcKochCurve(n-1, u, t)
	buffering(t)
	calcKochCurve(n-1, t, p2)
}

func main() {
	n := scanToInt()

	p1 := &node{
		x: 0,
		y: 0,
	}

	buffering(p1)

	p2 := &node{
		x: 100,
		y: 0,
	}

	calcKochCurve(n, p1, p2)
	buffering(p2)

	fmt.Println(buf.String())
}

var sc = bufio.NewScanner(os.Stdin)
var buf bytes.Buffer

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func buffering(n *node) {
	buf.WriteString(fmt.Sprintf("%.8f %.8f\n", n.x, n.y))
}
