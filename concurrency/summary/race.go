package summary

import (
	"fmt"
	"sync"
	"time"
)

var data int

// RaceUnsafe 竞争
// 多个操作同时都尝试去获取同一个内存的数据
func RaceUnsafe() {
	go func() {
		// 获取data的值 -> data+1 -> 把新data赋值
		// 并发不安全是因为其他操作在上面的操作可能还未结束，就对变量进行操作
		// 扩展：原子操作即操作不可再分割，不会同时发生别的事
		// 考虑原子性时，可能在某个上下文是原子性的，但是在另一个上下文又不是
		// data++ 在这里不是原子性的，但是如果没有并发，他是原子性的
		data++
	}()

	if data == 0 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(data)
	}

}

var memoryAccess sync.Mutex

// RaceSafeByLock 解决了数据竞争问题
// 但是没有真正解决竞争条件，因为顺序是不确定的
func RaceSafeByLock() {
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(data)
	}
	memoryAccess.Unlock()

}
