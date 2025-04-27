package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	var wg sync.WaitGroup
	// 创建一个无缓冲通道，发送和接收必须同时准备好
	ch := make(chan int)

	wg.Add(2)

	// Goroutine 1: 尝试从通道接收，但此时没有发送者，会阻塞
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1: Attempting to receive from channel...")
		// 这行会阻塞，直到 main Goroutine 发送数据
		val := <-ch
		fmt.Printf("Goroutine 1: Received value: %d\n", val)
	}()

	// Goroutine 2: 模拟做其他工作，证明 M 没有被阻塞
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2: Running concurrently...")
		time.Sleep(500 * time.Millisecond) // 做一些工作
		fmt.Println("Goroutine 2: Finished work.")
	}()

	// 主 Goroutine 等待一段时间，模拟 Goroutine 1 阻塞
	fmt.Println("Main Goroutine: Waiting for a second before sending...")
	time.Sleep(1 * time.Second)

	// 向通道发送数据，解除 Goroutine 1 的阻塞
	fmt.Println("Main Goroutine: Sending value 10 to channel...")
	ch <- 10
	fmt.Println("Main Goroutine: Value sent.")

	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("Main Goroutine: All goroutines finished.")
}
