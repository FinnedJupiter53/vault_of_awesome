package main

import "math/rand"

type character struct {
	Name  string
	Class string
	Race  string
	Str   int
	Dex   int
	Con   int
	Int   int
	Wis   int
	Cha   int
}

func newCharacter(name string) character {
	var c character
	c.Name = name
	c.Class = randomClass()
	c.Race = randomRace()
	c.Str = randomStat()
	c.Dex = randomStat()
	c.Con = randomStat()
	c.Int = randomStat()
	c.Wis = randomStat()
	c.Cha = randomStat()

	return c
}

func die(max int) int {
	return rand.Int()%max + 1
}

func randomRace() string {
	races := []string{
		"Human",
		"Dwarf",
		"Elf",
		"Half Orc",
	}

	chosenRace := rand.Int() % len(races)

	return races[chosenRace]
}

func randomClass() string {
	classes := []string{
		"Fighter",
		"Barbarian",
		"Druid",
		"Wizard",
	}

	chosenClass := rand.Int() % len(classes)

	return classes[chosenClass]
}

func randomStat() int {
	return die(6) + die(6) + die(6)
}
