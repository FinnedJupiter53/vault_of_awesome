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
	w.WriteString(fmt.Sprintf("Level: %v\n", char.Level))
	w.WriteString(fmt.Sprintf("Race: %v\n", char.Race))
	w.WriteString(fmt.Sprintf("Background: %v\n", char.Background))
	w.WriteString(fmt.Sprintf("Str: %v (%v)\n", char.Str.Value, char.Str.Modifier()))
	w.WriteString(fmt.Sprintf("Dex: %v (%v)\n", char.Dex.Value, char.Dex.Modifier()))
	w.WriteString(fmt.Sprintf("Con: %v (%v)\n", char.Con.Value, char.Con.Modifier()))
	w.WriteString(fmt.Sprintf("Int: %v (%v)\n", char.Int.Value, char.Int.Modifier()))
	w.WriteString(fmt.Sprintf("Wis: %v (%v)\n", char.Wis.Value, char.Wis.Modifier()))
	w.WriteString(fmt.Sprintf("Cha: %v (%v)\n", char.Cha.Value, char.Cha.Modifier()))
}
