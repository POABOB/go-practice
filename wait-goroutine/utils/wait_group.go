package utils

import (
	"sync"
	"time"
)

func WaitByWaitGroup(n int) {
	// 宣告 sync.WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		// 協程計數器++
		wg.Add(1)
		go func(i int) {
			// 協程計數器--
			defer wg.Done()
			time.Sleep(1 * time.Nanosecond)
			//fmt.Printf("Goroutine %d finished\n", i)
		}(i)
	}

	//fmt.Println("Waiting for all goroutines to finish")
	// 阻塞等待計數器為 0
	wg.Wait()
	//fmt.Println("All goroutines finished")
}
