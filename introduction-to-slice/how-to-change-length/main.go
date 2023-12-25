package main

import "fmt"

func main() {
	arr := [4]string{"a", "b", "c", "d"} // len=4, cap=4, value=[a b f d]

	s1 := arr[2:3] // len=1, cap=2, value=[c]
	fmt.Printf("slice s1: len=%v, cap=%v, value=%v\n", len(s1), cap(s1), s1)

	s1[0] = "f" // len=1, cap=2, value=[f]
	fmt.Printf("slice s1: len=%v, cap=%v, value=%v\n", len(s1), cap(s1), s1)

	// 忽略 start 或 end，代表從頭取 或是 取至尾端
	s2 := arr[:2] // len=2, cap=4, value=[a b]
	fmt.Printf("slice s2: len=%v, cap=%v, value=%v\n", len(s2), cap(s2), s2)

	s3 := arr[2:] // len=2, cap=2, value=[f d]
	fmt.Printf("slice s3: len=%v, cap=%v, value=%v\n", len(s3), cap(s3), s3)

	// 如果再使用 Slice 長度調整時，想要去調整容量，才會使用到 max
	s4 := arr[:2:3] // len=2, cap=3, value=[a b]
	fmt.Printf("slice s4: len=%v, cap=%v, value=%v\n", len(s4), cap(s4), s4)
}
