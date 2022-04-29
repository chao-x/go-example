package main

import (
	"fmt"
	"sync"
	"time"
)

var data int

func RaceUnsafe() {
	go func() {
		data++
	}()

	if data == 0 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(data)
	}

}

var memoryAccess sync.Mutex

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
