package summary

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func printSum(v1, v2 *value) {
	defer wg.Done()

	v1.mu.Lock()
	defer v1.mu.Unlock()

	// 这里极大增加了死锁的概率
	time.Sleep(2 * time.Second)

	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Printf("sum = %v\n", v1.value+v2.value)
}

// DeadLock 死锁
// 相互排斥，都拥有资源的独占权
// 等待条件，都拥有一个资源，并且等待额外资源
// 没有抢占，拥有资源只能等待他自己释放
// 循环等待，都在互相等待对方释放资源
func DeadLock() {
	var v1, v2 value
	wg.Add(2)
	go printSum(&v1, &v2)
	go printSum(&v2, &v1)
	wg.Wait()
}
