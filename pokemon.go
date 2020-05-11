package main

import "fmt"

type Pokemon struct {
	Name    string
	Element string
}

func (p *Pokemon) String() string {
	return fmt.Sprintf("Pokemon: %s, Type: %s", p.Name, p.Element)
}

func main() {

	fmt.Println("My Ten Favorite Pokemon and Their Types...")

	mewtwo := &Pokemon{
		Name:    "Mewtwo",
		Element: "Psychic",
	}

	alakazam := &Pokemon{
		Name:    "Alakazam",
		Element: "Psychic",
	}

	haunter := &Pokemon{
		Name:    "Haunter",
		Element: "Poison, Ghost",
	}

	charizard := &Pokemon{
		Name:    "Charizard",
		Element: "Fire, Flying",
	}

	gyarados := &Pokemon{
		Name:    "Gyarados",
		Element: "Water, Flying",
	}

	onix := &Pokemon{
		Name:    "Onix",
		Element: "Ground, Rock",
	}

	scyther := &Pokemon{
		Name:    "Scyther",
		Element: "Bug, Flying",
	}

	kabutops := &Pokemon{
		Name:    "Kabutops",
		Element: "Rock, Water",
	}

	zapdos := &Pokemon{
		Name:    "Zapdos",
		Element: "Electric, Flying",
	}

	ditto := &Pokemon{
		Name:    "Ditto",
		Element: "Normal",
	}

	fmt.Println(mewtwo)
	fmt.Println(alakazam)
	fmt.Println(haunter)
	fmt.Println(charizard)
	fmt.Println(gyarados)
	fmt.Println(onix)
	fmt.Println(scyther)
	fmt.Println(kabutops)
	fmt.Println(zapdos)
	fmt.Println(ditto)
}
