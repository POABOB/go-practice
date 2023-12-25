package main

import "fmt"

func main() {
	scores := []int{1, 2, 3, 4, 5}

	// cloneScores 容量大於 scores 複製過來的元素長度
	cloneScores := make([]int, 4)
	copy(cloneScores, scores[:len(scores)-2])
	fmt.Println(cloneScores) // [1,2,3,0]

	// cloneScores 容量小於 scores 複製過來的元素長度
	cloneScores2 := make([]int, 3)
	copy(cloneScores2, scores)
	fmt.Println(cloneScores2) // [1,2,3]
}
