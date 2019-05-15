package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	command = iota
	key
)

func main() {
	// scan
	n := scanToInt()

	buf := bytes.Buffer{}

	m := make(map[string]struct{}, n)

	for i := 0; i < n; i++ {
		s := strings.Split(scanToString(), " ")

		if s[command] == "insert" {
			m[s[key]] = struct{}{}
		} else {
			if _, ok := m[s[key]]; ok {
				buf.WriteString("yes\n")
			} else {
				buf.WriteString("no\n")
			}
		}
	}
	fmt.Print(buf.String())
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

func scanToString() string {
	sc.Scan()
	return sc.Text()
}
