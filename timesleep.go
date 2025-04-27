package main

import (
	"fmt"
	"sync"
	"time"
)

func main5() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: 调用 time.Sleep
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1: Starting to sleep for 2 seconds...")
		// 这个 Goroutine 会被调度器暂停，M 可以运行其他任务
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1: Woke up after sleep.")
	}()

	// Goroutine 2: 同时运行，证明 M 没有被 Goroutine 1 的 sleep 阻塞
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2: Running concurrently during sleep...")
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine 2: Working... %d\n", i+1)
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Goroutine 2: Finished work.")
	}()

	wg.Wait()
	fmt.Println("Main Goroutine: All goroutines finished.")
}
