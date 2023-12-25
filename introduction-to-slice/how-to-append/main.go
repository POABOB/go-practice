package main

import "fmt"

func main() {
	s := make([]int, 0, 1) // [] len=0 cap=1
	fmt.Printf("slice s: len=%v, cap=%v, value=%v\n", len(s), cap(s), s)
	s = append(s, 1) // [1] len=1 cap=1
	fmt.Printf("slice s: len=%v, cap=%v, value=%v\n", len(s), cap(s), s)
	s = append(s, 1) // [1 1] len=2 cap=2
	fmt.Printf("slice s: len=%v, cap=%v, value=%v\n", len(s), cap(s), s)
	s = append(s, 1) // [1 1 1] len=3 cap=4
	fmt.Printf("slice s: len=%v, cap=%v, value=%v\n", len(s), cap(s), s)

}
