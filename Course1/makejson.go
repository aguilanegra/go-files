package main

import (
	"bufio"
	"encoding/json"
	"fmt"
  	"os"
)

func main() {
	var name, address string

	// Ask user to enter name, use os.Stdin instead of Scan() to accommodate whitespaces
	fmt.Print("Enter a name: ")
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    name = scanner.Text()

	// Ask user to enter address, use os.Stdin instead of Scan() to accommodate whitespaces
	fmt.Print("Enter an address: ")
	scanner = bufio.NewScanner(os.Stdin)
    scanner.Scan()
    address = scanner.Text()

	// Create a map from the input data
	data := map[string]string{
		"name":    name,
		"address": address,
	}

	// Make JSON magic!
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output
	fmt.Println("JSON Object:", string(jsonData))
}
