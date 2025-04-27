package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main3() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: 发起一个网络请求，会 "阻塞" 等待响应
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1: Starting HTTP GET request...")
		// 这个调用看起来是阻塞的，但 M 不会阻塞
		resp, err := http.Get("https://httpbin.org/delay/2") // 请求一个会延迟 2 秒响应的 URL
		if err != nil {
			fmt.Println("Goroutine 1: Error:", err)
			return
		}
		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body) // 读取响应体
		if err != nil {
			fmt.Println("Goroutine 1: Error reading body:", err)
			return
		}
		fmt.Println("Goroutine 1: HTTP GET request finished.")
	}()

	// Goroutine 2: 同时运行，证明 M 没有被 Goroutine 1 的网络请求阻塞
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2: Running concurrently...")
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine 2: Working... %d\n", i+1)
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Goroutine 2: Finished work.")
	}()

	wg.Wait()
	fmt.Println("Main Goroutine: All goroutines finished.")
}
