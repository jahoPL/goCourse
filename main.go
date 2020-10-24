package main

import (
	"log"
)

func main() {

}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
