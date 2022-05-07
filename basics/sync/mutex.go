package _sync

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

// Mutex 首次加锁用时会比后面的长
// 随着goroutine数量增多，读写锁的性能会优于互斥锁
func Mutex() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(time.Nanosecond)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	// 格式化输出，用\t来识别分割
	// min width     最小单元长度
	// tab width     tab字符的宽度
	// padding       计算单元宽度时会额外加上它
	// padding char  用于填充的ASCII字符，
	//               如果是'\t'，则Writer会假设tab width作为输出中tab的宽度，且单元必然左对齐。
	// flags    格式化控制
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Reader\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		// RLocker()作用是，使用Lock()和Unlock()来进行读锁定和读解锁，而无需RLock()和RUnlock()来进行读锁定和读解锁
		fmt.Fprintf(tw, "%d\t%v\t%v\n", count, test(count, &m, m.RLocker()), test(count, &m, &m))
	}
}
