package main

import (
	"bufio"
	"fmt"
	"os"
)

// Name struct represents a person's first and last name
type Name struct {
	Fname string
	Lname string
}

func main() {
	var filename string

	// Prompt the user for the name of the text file
	fmt.Print("Enter the name of the text file: ")
	fmt.Scan(&filename)

	// Open the text file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a slice to store the structs
	names := make([]Name, 0)

	/* 
	Read each line from the file.
	Assumptions: The text file conatains a listing of first and last names separated
	by a space on each line. First name and last name fields are respectively no longer
	than 20 characters, so the max length of any single line including the space is 41.
	*/
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into first name and last name
		namesArr := splitName(line)

		// Create a new Name struct and add it to the slice
		name := Name{
			Fname: namesArr[0],
			Lname: namesArr[1],
		}
		names = append(names, name)
	}

	// Print the first and last names from each struct
	for _, name := range names {
		fmt.Println("\nFirst Name:", name.Fname)
		fmt.Println("Last Name:", name.Lname)
	}
}

// splitName splits a line into first name and last name
func splitName(line string) [2]string {
	namesArr := [2]string{"", ""}
	for i, name := range line {
		if name == ' ' {
			namesArr[0] = line[:i]
			namesArr[1] = line[i+1:]
			break
		}
	}
	return namesArr
}