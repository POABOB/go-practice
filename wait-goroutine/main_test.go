package main

import (
	"github.com/POABOB/go-practice/wait-goroutine/utils"
	"testing"
)

func BenchmarkWaitByChannelCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.WaitByChannelCount(100)
	}
}

func BenchmarkWaitByWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.WaitByWaitGroup(100)
	}
}

func BenchmarkSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.Sleep(100)
	}
}
