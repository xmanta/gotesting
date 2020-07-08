package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	 rwlock.Lock()
	 x = x +1
	 time.Sleep(10 * time.Microsecond)
	 rwlock.Unlock()
	 wg.Done()
}

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i <10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i <10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}