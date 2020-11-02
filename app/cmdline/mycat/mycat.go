package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		debug bool
	)
	flag.BoolVar(&debug, "n", false, "desplay line number")
	flag.Parse()

	args := os.Args

	if len(args) == 1 {
		log.Fatal("ファイル名を引数にいれて")
	}
	for _, arg := range args[1:] {
		if debug {
			readn(arg)
		} else {
			read(arg)
		}
	}
}

func read(file string) error {
	sf, err := os.Open(file)
	if err != nil {
		return err
	}
	// 関数終了時に閉じる
	defer sf.Close()
	scanner := bufio.NewScanner(sf)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func readn(file string) error {
	sf, err := os.Open(file)
	if err != nil {
		return err
	}
	// 関数終了時に閉じる
	defer sf.Close()
	scanner := bufio.NewScanner(sf)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(i, " ", line)
		i++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
