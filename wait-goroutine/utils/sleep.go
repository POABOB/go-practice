package utils

import (
	"time"
)

func Sleep(n int) {
	for i := 0; i < n; i++ {
		go func(i int) {
			time.Sleep(1 * time.Nanosecond)
			//fmt.Printf("Goroutine %d finished\n", i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	//fmt.Println("All goroutines finished")
}
