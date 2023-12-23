package utils

import "time"

func WaitByChannelCount(n int) {
	done := make(chan bool)
	for i := 0; i < n; i++ {
		go func(done chan bool) {
			time.Sleep(1 * time.Nanosecond)
			//fmt.Println("working")
			done <- true
		}(done)
	}

	//fmt.Println("waiting")

	i := 0
Loop:
	for {
		select {
		case <-done:
			i++
		default:
			if i == n {
				break Loop
			}
		}
	}
	//fmt.Println("done")
}
