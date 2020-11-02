package main

import "errors"

func main() {
	if err := f(); err != nil {
		// error handle
	}
}
func f() error {
	return errors.New("hoge")
}
