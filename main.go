package main

import (
	"github/Ko4s/goCourse/topic3"
	"log"
)

func main() {
	h := topic3.NewHasher("topic3/in.txt", "topic3/out.txt")
	err := h.ReadLines()
	handleError(err)

	err = h.HashLines()
	handleError(err)

	err = h.Save()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
