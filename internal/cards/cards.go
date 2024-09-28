package cards

import "fmt"

type Card struct {
	Name     string
	Strength int
	Agility  int
	Defense  int
}

func (c *Card) MountCardInterface() {
	fmt.Println("Nome: ", c.Name)
	fmt.Println("-----------------")
	fmt.Println("0 - Força", c.Strength)
	fmt.Println("1 - Defesa", c.Defense)
	fmt.Println("2 - Agilidade", c.Agility)
}

type Attribute int

var AttributeMap = map[Attribute]string{
	0: "Strength",
	1: "Defense",
	2: "Agility",
}

var Gods = []Card{
	{
		Name:     "Zeus",
		Strength: 10,
		Agility:  5,
		Defense:  7,
	},
	{
		Name:     "Poseidon",
		Strength: 8,
		Agility:  7,
		Defense:  9,
	},
	{
		Name:     "Hades",
		Strength: 9,
		Agility:  6,
		Defense:  8,
	},
	{
		Name:     "Athena",
		Strength: 7,
		Agility:  9,
		Defense:  6,
	},
	{
		Name:     "Ares",
		Strength: 10,
		Agility:  4,
		Defense:  8,
	},
	{
		Name:     "Apollo",
		Strength: 6,
		Agility:  8,
		Defense:  5,
	},
	{
		Name:     "Artemis",
		Strength: 7,
		Agility:  10,
		Defense:  5,
	},
	{
		Name:     "Hera",
		Strength: 6,
		Agility:  7,
		Defense:  8,
	},
	{
		Name:     "Hermes",
		Strength: 5,
		Agility:  10,
		Defense:  6,
	},
	{
		Name:     "Hephaestus",
		Strength: 9,
		Agility:  4,
		Defense:  10,
	},
}
