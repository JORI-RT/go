package main

import (
	"flag"
	"fmt"
	"strings"
)

var msg = flag.String("msg", "デフォルト値", "説明")
var n = flag.Int("n", 1, "回数")

func init() {
	// ポインタを指定して設定を予約
}
func main() {
	// ここで実際に設定される
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, *n))
}
