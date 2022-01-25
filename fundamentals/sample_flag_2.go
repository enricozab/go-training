package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("i", "input.dat", "text or binary file")
	flag.Parse()
	data, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatalf("Cannot read file %v: %v\n", *filename, err)
	}

	fmt.Println("%x\n", sha256.Sum256(data))
}
