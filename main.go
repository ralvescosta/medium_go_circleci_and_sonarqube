package main

import (
	"log"

	"ci/pkg"
)

func main() {
	log.Printf("1 + 1 = %v\n", pkg.Sum(1, 1))
	log.Printf("ABS(-2) = %v\n", pkg.Abs(-2))
}
