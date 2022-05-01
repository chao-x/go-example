package summary

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func LiveLock() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		// 每隔一毫秒执行一次
		for range time.Tick(1 * time.Millisecond) {
			// 唤醒 Wait的协程
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		// Wait 做了啥？
		// 解锁 -> 切走goroutine -> 等待唤醒 -> 上锁
		// 这里解锁才能使其他线程获得锁
		cadence.Wait()
		// 这里必须解锁，因为Wait又上锁了
		cadence.L.Unlock()
	}

	tryDire := func(direName string, dire *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", direName)
		atomic.AddInt32(dire, 1)
		// takeStep 会切换协程，所以导致了活锁的发生
		takeStep()
		if atomic.LoadInt32(dire) == 1 {
			fmt.Fprintf(out, ". Success!")
			return true
		}
		takeStep()
		atomic.AddInt32(dire, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDire("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDire("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer walking.Done()
		defer func() { fmt.Println(out.String()) }()
		fmt.Fprintf(&out, "%v is trying to scoot：", name)

		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
