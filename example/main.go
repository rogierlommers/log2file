package main

import (
	"log"

	"github.com/rogierlommers/log2file"
)

func init() {
	if err := log2file.Activate(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	for i := 0; i <= 5; i++ {
		log2file.Write("line")
	}
}
