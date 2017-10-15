package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func die(max int) int {
	return rand.Int()%max + 1
}

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	num := flag.Int("num", 1, "number of dice")
	max := flag.Int("max", 20, "max value of dice")

	if *num <= 0 {
		fmt.Println("'num' flag must have a number greater than zero")
	}
	if *max <= 0 {
		fmt.Println("'max' flag must have a number greater than zero")
	}

	for i := 0; i < *num; i++ {
		fmt.Printf("%v\n", die(*max))
	}
}
