package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	name := flag.String("name", "Boros", "name of the character")

	flag.Parse()

	char := newCharacter(*name)

	fileName := fmt.Sprintf("%s.txt", *name)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	prettyPrintFile(char, f)
}

func prettyPrintFile(char character, w *os.File) {
	w.WriteString(fmt.Sprintf("Name: %v\n", char.Name))
	w.WriteString(fmt.Sprintf("Class: %v\n", char.Class))
	w.WriteString(fmt.Sprintf("Race: %v\n", char.Race))
	w.WriteString(fmt.Sprintf("Str: %v\n", char.Str))
	w.WriteString(fmt.Sprintf("Dex: %v\n", char.Dex))
	w.WriteString(fmt.Sprintf("Con: %v\n", char.Con))
	w.WriteString(fmt.Sprintf("Int: %v\n", char.Int))
	w.WriteString(fmt.Sprintf("Wis: %v\n", char.Wis))
	w.WriteString(fmt.Sprintf("Cha: %v\n", char.Cha))
}
