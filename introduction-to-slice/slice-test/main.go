package main

import "fmt"

func main() {
	// Test 1
	fmt.Println("Test 1")
	var array [20]int
	var slice = array[10:11]
	fmt.Println("length: ", len(slice))
	fmt.Println("capacity: ", cap(slice))
	fmt.Println(&slice[0] == &array[10])

	fmt.Println()
	// Test 2
	fmt.Println("Test 2")
	orderLen := 5
	order := make([]uint16, 2*orderLen)

	pollOrder := order[:orderLen:orderLen]
	lockOrder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollOrder) = ", len(pollOrder))
	fmt.Println("cap(pollOrder) = ", cap(pollOrder))
	fmt.Println("len(lockOrder) = ", len(lockOrder))
	fmt.Println("cap(lockOrder) = ", cap(lockOrder))

	fmt.Println()
	// Test 3
	fmt.Println("Test 3")
	s1 := []int{1, 2}
	s1 = append(s1, 3, 4, 5)
	fmt.Printf("cap of s1:%d\n", cap(s1))

	s2 := make([]int, 255, 255)
	s2 = append(s2, 1)
	fmt.Printf("cap of s2:%d\n", cap(s2))

	s3 := make([]int, 257, 257)
	s3 = append(s3, 1)
	fmt.Printf("cap of s3:%d\n", cap(s3))

	s4 := make([]int, 512, 512)
	s4 = append(s4, 1)
	fmt.Printf("cap of s4:%d\n", cap(s4))
}
