package main

import (
	"fmt"
	"github.com/POABOB/go-practice/wait-goroutine/utils"
)

func main() {
	go func() {
		fmt.Println("say hello......")
	}()
	fmt.Println("main groutine....")

	utils.WaitByChannelCount(100)

	utils.WaitByWaitGroup(100)

	utils.Sleep(100)
}
