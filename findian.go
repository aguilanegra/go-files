package main

import (
	"fmt"
	"bufio"
  	"os"
	"strings"
	"regexp"
)

func main() {
	fmt.Print("Enter a string: ")
	// Use os.Stdin instead of Scan() to accommodate whitespaces
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    line := scanner.Text()

	// Create a regular expression pattern
	pattern := "(?i)^\\s*i.*a.*n\\s*$"

	// Remove all whitespace characters from the input
	processedInput := strings.ReplaceAll(line, " ", "")
	matched, _ := regexp.MatchString(pattern, processedInput)

	if matched {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}