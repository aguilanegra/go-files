package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3) // Create an empty integer slice of size 3

	for {
		var input string
		fmt.Print("Enter an integer (or 'X' to exit): ")
		fmt.Scan(&input)

		// Entering "X" will exit the loop and end the program
		if input == "X" {
			break
		}

		// If you are really trying to break me ;)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or 'X'.")
			fmt.Println("Sorted Slice:", slice)
			continue
		}

		// This is where the good stuff is, sort it!
		slice = append(slice, num)
		sort.Ints(slice)

		fmt.Println("Sorted Slice:", slice)
	}
}
