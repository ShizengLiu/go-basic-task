package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func demo1() {

	go func() {
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数：", i)
		}
	}()

	go func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数：", i)
		}
	}()

}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用线程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：线程原理、并发任务调度。

type Task struct {
	id int
}

func (t *Task) run() {
	startTime := time.Now()
	sleepDuration := time.Duration(rand.Intn(2)+1) * time.Second
	time.Sleep(sleepDuration)
	fmt.Println("task run", t.id, "task time", time.Since(startTime))
}

func taskDemo() {
	var wg sync.WaitGroup

	workers := make([]Task, 0)
	for i := 0; i < 5; i++ {
		task := Task{id: i}
		workers = append(workers, task)
	}
	wg.Add(len(workers))
	for i := 0; i < len(workers); i++ {
		go func(taskId int) {
			defer wg.Done()
			workers[taskId].run()
		}(i)
	}

	wg.Wait()

}

func main() {
	demo1()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("--------------------------")
	taskDemo()
}
