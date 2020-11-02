package main

import (
	"flag"
	"fmt"
	"strings"
)

// --help でフラグの使い方が見える
var msg = flag.String("msg", "デフォルト値", "説明")
var n = flag.Int("n", 1, "回数")

func init() {
	// ポインタを指定して設定を予約
}
func main() {
	var (
		debug bool
	)
	flag.BoolVar(&debug, "debug", false, "Debug Mode enabled?")
	// ここで実際に設定される
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, *n))
	fmt.Println(flag.Args())
	fmt.Printf("param -debug -> %t\n", debug)

}
