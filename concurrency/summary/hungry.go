package summary

import (
	"fmt"
	"sync"
	"time"
)

// Hungry 并发进程所需资源被其他的并发进程一直占有
// 下面的例子中，贪婪和平和的两个进程的锁力度是不一样的
// 尽管允许结果贪婪进程做了更多 但是并不表示贪婪进程的效率更好
// 而是因为贪婪进程（可能）在不需要锁的时候占有锁，导致平和进程经常处于“饥饿”状态
func Hungry() {
	var wg sync.WaitGroup
	var shareLock sync.Mutex
	const runtime = 1*time.Second

	// greedy 贪婪的获取锁
	greedy := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) < runtime; {
			shareLock.Lock()
			time.Sleep(3*time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops.\n", count)

	}

	// politeWork 只在需要的时候获取锁
	politeWork := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) < runtime; {
			shareLock.Lock()
			time.Sleep(time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(time.Nanosecond)
			shareLock.Unlock()

			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}


	wg.Add(2)
	go politeWork()
	go greedy()
	wg.Wait()
}