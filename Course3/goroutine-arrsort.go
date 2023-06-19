package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
    // Prompt user for the integers to sort
	fmt.Println("Enter a series of integers (space-separated):")

    // Trying to be fault-tolerant with the input data by trimming
    // spaces from the input data
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numbers := strings.Split(input, " ")

    // Split the work into 4 evenly-divided partitions, and if
    // the number of integers entered is less than four then
    // evenly divide the partitions
	partitions := 4
	if len(numbers) < partitions {
		partitions = len(numbers)
	}

	var wg sync.WaitGroup
	ch := make(chan []int, partitions)

	partitionSize := len(numbers) / partitions
	extraElements := len(numbers) % partitions

    // Each partition sorts its respective subarray within a buffered channel
	start := 0
	for i := 0; i < partitions; i++ {
		end := start + partitionSize
		if i < extraElements {
			end++
		}

		subarray := numbers[start:end]
		wg.Add(1)
		go sortSubarray(subarray, ch, &wg)

		start = end
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var sorted []int
	for subarray := range ch {
		sorted = mergeSortedArrays(sorted, subarray)
	}

    // Print the final sorted main array
	fmt.Println("Sorted array:")
	fmt.Println(sorted)
}

// Sorting function for subarray
func sortSubarray(subarray []string, ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()

    ints := make([]int, len(subarray))
    for i, numStr := range subarray {
        ints[i], _ = strconv.Atoi(numStr)
    }

    // Print out the array to be sorted, per the instructions
    fmt.Println("Subarray to sort:", ints)

    sort.Ints(ints)

    // Print the sorted subarray
    //fmt.Println("Sorted subarray:", ints)
    ch <- ints
}

// Function to merge the respective subarrays into the main array
func mergeSortedArrays(a, b []int) []int {
	merged := make([]int, 0, len(a)+len(b))
    i, j := 0, 0

    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            merged = append(merged, a[i])
            i++
        } else {
            merged = append(merged, b[j])
            j++
        }
    }

    merged = append(merged, a[i:]...)
    merged = append(merged, b[j:]...)

	return merged
}
