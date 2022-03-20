package main

import (
	"log"

	"ci/pkg"
)

func main() {
	log.Println("running...")
	log.Printf("1 + 1 = %v\n", pkg.Sum(1, 1))
	log.Printf("ABS(-2) = %v\n", pkg.Abs(-2))
	log.Println("finished")
}
