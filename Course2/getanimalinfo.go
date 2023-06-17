package main

import "fmt"

// Animal represents an animal with its food, locomotion, and noise characteristics
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints the animal's food
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints the animal's locomotion
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak prints the animal's spoken sound
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
    // map data structure holding animal data
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	for {
		var animalName, info string

		// Prompt the user for an animal name and information request
		fmt.Print("> ")
		fmt.Scanln(&animalName, &info)

		// Get the animal from the map
		animal, found := animals[animalName]
		if !found {
			fmt.Println("Invalid animal name")
			continue
		}

		// Process the information request
		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid information request")
		}
	}
}