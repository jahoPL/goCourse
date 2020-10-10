package main

import (
	"github/Ko4s/goCourse/topic3"
	"log"
)

func main() {
	h := topic3.NewHasher("in","out")

	err := h.ReadLines()
	handleError(err)

	err = h.HashLines()
	handleError(err)

	err = h.SaveToFile()
	handleError(err)
}
func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
