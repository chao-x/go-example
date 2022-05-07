package component

import (
	"fmt"
	"runtime"
	"sync"
)

// FatherGoroutine main goroutine退出所有goroutine都会退出
// 但是其他情况下，父goroutine退出，子goroutine并不会退出
// 所以有可能会导致goroutine泄露的内存泄露问题
func FatherGoroutine(wg *sync.WaitGroup) {
	sonGoroutine := func() {
		defer wg.Done()
		for i := 100; i < 200; i++ {
			fmt.Println(i)
		}
	}
	go sonGoroutine()
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}

// FatherGoroutineMain FatherGoroutine的main方法示例
// FatherGoroutine结束后，他的子goroutine：sonGoroutine仍会进行
func FatherGoroutineMain() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	FatherGoroutine(wg)
	wg.Wait()
}

// Address 闭包打印的salutation用的是同一个地址
// 大概率在goroutine还没开始前循环就已经结束
// 因为goroutine开始需要准备时间
func Address() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation) // 打印3个good day
		}()
	}
	wg.Wait()
}

// AddressSalutation 防止重复打印的解决方法
// 将参数copy传入,但三个字符串不会顺序执行
func AddressSalutation() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation) // 打印3个good day
		}(salutation)
	}
	wg.Wait()
}

// Memory goroutine
//子routine都是空的，但是不退出
// 10**4 个空goroutine 平均每个所占用的内存
func Memory() {
	// 当前程序所占用内存
	memConsumed := func() uint64 {
		// 显式的触发GC
		// 需要手动强制触发的场景极其少见，可能会是在某些业务方法执行完后，因其占用了过多的内存，需要人为释放。又或是debug程序所需
		runtime.GC()
		var s runtime.MemStats
		// 已分配内存的总量，单位是 b
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	// noop永远被c阻塞
	noop := func() { wg.Done(); <-c }

	// goroutine数量
	const numGoroutine = 1e4
	wg.Add(numGoroutine)
	before := memConsumed()
	for i := numGoroutine; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutine/1024)
}
