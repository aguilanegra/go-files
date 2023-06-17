package main

import (
	"fmt"
)

// Animal interface describes the methods of an animal
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow stuff //
// Cow represents a cow
type Cow struct {
	name string
}

// Eat prints the food of the cow
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move prints the locomotion method of the cow
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak prints the sound made by the cow
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird stuff //
// Bird represents a bird
type Bird struct {
	name string
}

// Eat prints the food of the bird
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move prints the locomotion method of the bird
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak prints the sound made by the bird
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake stuff //
// Snake represents a snake
type Snake struct {
	name string
}

// Eat prints the food of the snake
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move prints the locomotion method of the snake
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak prints the sound made by the snake
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		var command, name, action string

		// Prompt the user for a command
		fmt.Print("> ")
		fmt.Scanln(&command, &name, &action)

		switch command {
		case "newanimal":
			// Create a new animal
			var newAnimal Animal

			switch action {
			case "cow":
				newAnimal = Cow{name}
			case "bird":
				newAnimal = Bird{name}
			case "snake":
				newAnimal = Snake{name}
			default:
				fmt.Println("Invalid animal type")
				continue
			}

			// Add the new animal to the map
			animals[name] = newAnimal

			fmt.Println("Created it!")
		case "query":
			// Retrieve information about an animal
			animal, found := animals[name]
			if !found {
				fmt.Println("Animal not found")
				continue
			}

			// Process the information request
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid information request")
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}
