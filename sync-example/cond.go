package sync_example

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Cond() {
	wg := sync.WaitGroup{}
	cond := sync.NewCond(new(sync.Mutex))
	for i := 0; i < 3; i++ {
		go func(i int) {
			wg.Add(1)
			fmt.Println("协程启动" + strconv.Itoa(i))

			cond.L.Lock()
			fmt.Println("协程  ", i, "加锁")

			// Wait 做了什么
			// 解锁 -> 切走goroutine -> 等待唤醒 -> 上锁
			// 这里解锁才能使其他线程获得锁
			cond.Wait()

			cond.L.Unlock()
			fmt.Println("协程  ", i, "解锁")

			defer wg.Done()
		}(i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("主协程发送信号")
	cond.Signal()

	time.Sleep(time.Second * 2)
	fmt.Println("主协程发送信号")
	cond.Signal()

	time.Sleep(time.Second * 2)
	fmt.Println("主协程发送信号")
	cond.Signal()

	wg.Wait()
}
