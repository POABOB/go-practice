package main

import "fmt"

func main() {
	// 法一 make([]type, len, cap)
	// 宣告一個存放 string 的 Slice，長度和容量都會是 4
	s := make([]string, 4) // len=4, cap=4, value=["", "", "", ""]
	fmt.Printf("slice s: len=%v, cap=%v, value=%v\n", len(s), cap(s), s)

	// 長度 0，容量是 8
	s1 := make([]string, 0, 8) // len=0, cap=8, value=[]
	fmt.Printf("slice s1: len=%v, cap=%v, value=%v\n", len(s1), cap(s1), s1)

	// 法二 定義明確資料
	s2 := []string{"Alice", "Bob"} // len=2, cap=2, value=["Alice" "Bob"]
	fmt.Printf("slice s2: len=%v, cap=%v, value=%v\n", len(s2), cap(s2), s2)

	// 法三 宣告空 Slice
	var s3 []string // len=0, cap=0, value=[]
	fmt.Printf("slice s3: len=%v, cap=%v, value=%v\n", len(s3), cap(s3), s3)
}
