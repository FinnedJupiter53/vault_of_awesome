package main

import "math/rand"

type character struct {
	attributes
	Name       string
	Level      int
	Class      string
	Race       string
	Background string
}

func newCharacter(name string) character {
	var c character
	c.Name = name
	c.Level = 1
	c.Class = randomClass()
	c.Race = randomRace()
	c.Background = randomBackground()
	c.attributes = newAttributes()

	return c
}

func die(max int) int {
	return rand.Int()%max + 1
}

func randomRace() string {
	races := []string{
		"Human (Calshite)",
		"Human (Chondathan)",
		"Human (Damaran)",
		"Human (Illuskan)",
		"Human (Mulan)",
		"Human (Rashimi)",
		"Human (Shou)",
		"Human (Tethyrian)",
		"Human (Turami)",
		"Dwarf",
		"Mountain Dwarf",
		"Duergar",
		"High Elf",
		"Wood Elf",
		"Dark Elf",
		"Half-Orc",
		"Halfling",
		"Dragonborn",
		"Gnome",
		"Forest Gnome",
		"Rock Gnome",
		"Half-Elf",
		"Tiefling",
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
		"Bard",
		"Cleric",
		"Monk",
		"Ranger",
		"Rogue",
		"Warlock",
	}

	chosenClass := rand.Int() % len(classes)

	return classes[chosenClass]
}

func randomBackground() string {
	backgrounds := []string{
		"Acolyte",
		"Charlatan",
		"Criminal",
		"Entertainer",
		"Folk Hero",
		"Guild Artisan",
		"Hermit",
		"Noble",
		"Outlander",
		"Sage",
		"Sailor",
		"Soldier",
		"Urchin",
	}
	chosenBackground := rand.Int() % len(backgrounds)

	return backgrounds[chosenBackground]
}
