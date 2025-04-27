package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main4() {
	runtime.GOMAXPROCS(1) // 限制为 1 个 P，更容易观察 M 的行为（非必需，但有助于理解）
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Printf("Initial number of OS threads (approx): %d\n", runtime.NumGoroutine()) // 初始 Goroutine 数，间接反映线程

	// Goroutine 1: 执行一个阻塞的系统调用 (读取标准输入)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1: Waiting for stdin input (press Enter)...")
		// ReadString 会进行阻塞的系统调用，无法被 NetPoller 处理
		reader := bufio.NewReader(os.Stdin)
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Goroutine 1: Error reading stdin:", err)
			return
		}
		fmt.Println("Goroutine 1: Received stdin input.")
	}()

	// Goroutine 2: 模拟工作，观察是否能在 Goroutine 1 阻塞时运行
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2: Starting concurrent work...")
		// 打印线程数，可能观察到运行时为了让 Goroutine 2 运行而增加了 M
		// 注意：NumGoroutine 不是直接线程数，但间接相关；更精确的观察需要调试工具
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine 2: Working... %d\n", i+1)
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Goroutine 2: Finished work.")
	}()

	// 等待 Goroutine 启动
	time.Sleep(100 * time.Millisecond)

	// 提示用户输入
	fmt.Println("Main Goroutine: Program is waiting for input in Goroutine 1.")
	fmt.Println("Main Goroutine: Observe if Goroutine 2 continues to run.")

	wg.Wait()
	fmt.Println("Main Goroutine: All goroutines finished.")
}
