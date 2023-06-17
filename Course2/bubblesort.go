package main

import (
    "bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BubbleSort sorts the given slice of integers using the Bubble Sort algorithm
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

// Swap swaps the position of two adjacent elements in the slice
func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func main() {
	var input string
	var arr []int

	// Prompt the user to enter a sequence of up to 10 integers
	// use bufio scanner to handle whitespace characters
	fmt.Println("Enter a sequence of up to 10 integers (space-separated):")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input = scanner.Text()

	// Split the input string into individual integers
	strArr := strings.Fields(input)
	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invalid input:", str)
			return
		}
		arr = append(arr, num)
	}

	// Check if the number of elements exceeds 10
	if len(arr) > 10 {
		fmt.Println("You entered more than 10 integers, please run again and enter up to 10 integers.")
		return
	}

	// Sort the array using Bubble Sort
	BubbleSort(arr)

	// Print the sorted array
	fmt.Println("Sorted sequence:", arr)
}