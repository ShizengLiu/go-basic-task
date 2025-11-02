package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

func safeAdd() {
	var mutex sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	workers := 10
	incrementsTarget := 1000

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < incrementsTarget; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}

		}(i)
	}

	wg.Wait()

	fmt.Println("counter:", counter)
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

func atomicAdd() {
	var counter int32
	var wg sync.WaitGroup

	workers := 10
	incrementsTarget := 1000

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < incrementsTarget; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("atomic counter:", counter)

}

func main() {
	safeAdd()
	atomicAdd()
}
