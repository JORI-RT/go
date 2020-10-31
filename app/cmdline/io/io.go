package main

import (
	"errors"
	"log"
)

// https://docs.google.com/presentation/d/1Ui6zPU0bVuzmJLmlQUatxou0G93CdYHjm2nGPZrSC_Q/edit#slide=id.g8b2678b85b_0_705
func main() {
	// fmt.Fprintf(os.Stderr, "エラー")
	// os.Exit(1)
	if err := f(); err != nil {
		log.Fatal(err)
	}
}
func f() error {
	return errors.New("(*>△<)<ナーンナーンっっ")
}
