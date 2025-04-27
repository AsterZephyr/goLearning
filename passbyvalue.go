package main

import "fmt"

type MyStruct struct {
	Value int
}

func modifyData(
	num int, // 值类型
	str string, // 值类型
	s MyStruct, // 值类型 (结构体)
	ptrS *MyStruct, // 指针类型
	sl []int, // 切片 "引用类型"
	mp map[string]int, // Map "引用类型"
) {
	fmt.Println("\n--- Inside modifyData ---")

	// 1. 修改值类型 (int, string, struct) - 不影响外部
	num = 100
	str = "Modified String"
	s.Value = 200
	fmt.Printf("Inside: num=%d, str='%s', s.Value=%d\n", num, str, s.Value)

	// 2. 修改指针指向的结构体 - 影响外部
	if ptrS != nil {
		ptrS.Value = 300 // 通过指针副本修改原始数据
		fmt.Printf("Inside: ptrS.Value=%d (modified via pointer)\n", ptrS.Value)
	}

	// 3. 修改切片元素 - 影响外部 (因为共享底层数组)
	if len(sl) > 0 {
		sl[0] = 999
		fmt.Printf("Inside: sl[0]=%d (modified element)\n", sl[0])
	}

	// 4. 对切片进行 append (可能影响/不影响外部，取决于是否扩容)
	sl = append(sl, 4) // 如果容量足够，修改共享数组；如果容量不够，会分配新数组，副本sl指向新数组
	sl = append(sl, 5) // 进一步append
	fmt.Printf("Inside: sl after append=%v (might point to new array)\n", sl)

	// 5. 修改 Map 元素 - 影响外部 (因为共享底层哈希表)
	if mp != nil {
		mp["b"] = 222
		mp["c"] = 333
		fmt.Printf("Inside: mp after modification=%v\n", mp)
	}

	// 6. 给 Map 参数赋 nil - 不影响外部
	mp = nil
	fmt.Printf("Inside: mp after assigning nil=%v\n", mp)

	fmt.Println("--- Exiting modifyData ---")
}

func main1() {
	// 原始数据
	num := 10
	str := "Original String"
	s := MyStruct{Value: 20}
	ptrS := &MyStruct{Value: 30}
	sl := []int{1, 2, 3}
	originalSliceCapacity := cap(sl) // 记录原始容量
	mp := map[string]int{"a": 1, "b": 2}

	fmt.Println("--- Before modifyData ---")
	fmt.Printf("Outside: num=%d, str='%s', s.Value=%d\n", num, str, s.Value)
	fmt.Printf("Outside: ptrS.Value=%d\n", ptrS.Value)
	fmt.Printf("Outside: sl=%v (cap=%d)\n", sl, originalSliceCapacity)
	fmt.Printf("Outside: mp=%v\n", mp)

	// 调用函数，传递参数 (都是值传递)
	modifyData(num, str, s, ptrS, sl, mp)

	fmt.Println("\n--- After modifyData ---")
	fmt.Printf("Outside: num=%d (Unchanged)\n", num)
	fmt.Printf("Outside: str='%s' (Unchanged)\n", str)
	fmt.Printf("Outside: s.Value=%d (Unchanged)\n", s.Value)
	fmt.Printf("Outside: ptrS.Value=%d (Changed via pointer)\n", ptrS.Value)
	// 注意：这里 sl 的元素可能被修改，但 append 的结果是否可见取决于容量
	fmt.Printf("Outside: sl=%v (Element [0] changed, appends might be lost if reallocated)\n", sl)
	fmt.Printf("Outside: mp=%v (Changed, nil assignment inside function had no effect)\n", mp)

	// 检查 slice 容量是否改变，可以帮助判断 append 是否影响了外部
	fmt.Printf("Outside: sl capacity now %d (was %d)\n", cap(sl), originalSliceCapacity)
	if cap(sl) > originalSliceCapacity {
		fmt.Println("Slice was reallocated during append inside the function.")
	} else {
		fmt.Println("Slice was not reallocated, appends modified the shared underlying array (up to original capacity).")
	}
}
