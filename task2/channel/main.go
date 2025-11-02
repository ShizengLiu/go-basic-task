package main

import (
	"fmt"
	"time"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println("receive:", v)
	}
}

func send(ch chan<- int, area int) {
	for i := 0; i <= area; i++ {
		ch <- i
		fmt.Println("send:", i)
	}
	close(ch)
}

func demo1() {
	ch := make(chan int)
	go send(ch, 10)
	go receive(ch)
	time.Sleep(10 * time.Second)
}

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

func demo2() {
	ch := make(chan int, 50)
	go send(ch, 100)
	go receive(ch)
	time.Sleep(10 * time.Second)

}

func main() {
	demo1()
	fmt.Println("-----------------")
	demo2()
}
