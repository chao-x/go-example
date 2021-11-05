package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu sync.Mutex
	value int
}

var wg sync.WaitGroup

func printSum(v1,v2 *value) {
	defer wg.Done()

	v1.mu.Lock()
	defer v1.mu.Unlock()

	time.Sleep(2 * time.Second)

	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Printf("sum = %v\n", v1.value+v2.value)
}

func main() {
	var v1,v2 value
	wg.Add(2)
	go printSum(&v1, &v2)
	go printSum(&v2, &v1)
	wg.Wait()
}
