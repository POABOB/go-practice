package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8}

	for key, value := range pow {
		fmt.Printf("2**%d = %d\n", key, value)
	}

	for _, value := range pow {
		fmt.Printf("value = %v\n", value)
	}

	for key := range pow {
		fmt.Printf("key = %v\n", key)
	}
}
