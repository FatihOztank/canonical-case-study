package main

import (
	shred "example/exercise2/shred"
	"log"
)

func main() {
	// copy the file first, copied file will be deleted by the shred function
	err := shred.CopyFile("input.txt", "text.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = shred.Shred("text.txt")
	if err != nil {
		log.Fatal(err)
	}

}
