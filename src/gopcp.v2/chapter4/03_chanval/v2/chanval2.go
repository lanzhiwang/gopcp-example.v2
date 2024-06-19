package main

import (
	"fmt"
	"time"
)

// Counter 代表计数器的类型。
type Counter struct {
	count int
}

var mapChan = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { // 用于演示接收操作。
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
				fmt.Printf("The counter: %#v. [receiver]\n", counter)
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()
	go func() { // 用于演示发送操作。
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %#v. [sender]\n", countMap["count"])
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

// $ ./main
// The counter: main.Counter{count:1}. [receiver]
// The count map: map[string]main.Counter{"count":main.Counter{count:0}}. [sender]
// The counter: main.Counter{count:1}. [receiver]
// The count map: map[string]main.Counter{"count":main.Counter{count:0}}. [sender]
// The counter: main.Counter{count:1}. [receiver]
// The count map: map[string]main.Counter{"count":main.Counter{count:0}}. [sender]
// The counter: main.Counter{count:1}. [receiver]
// The count map: map[string]main.Counter{"count":main.Counter{count:0}}. [sender]
// The counter: main.Counter{count:1}. [receiver]
// The count map: map[string]main.Counter{"count":main.Counter{count:0}}. [sender]
// Stopped. [receiver]
// $

// $ ./main
// The counter: &main.Counter{count:1}. [receiver]
// The count map: &main.Counter{count:1}. [sender]
// The counter: &main.Counter{count:2}. [receiver]
// The count map: &main.Counter{count:2}. [sender]
// The counter: &main.Counter{count:3}. [receiver]
// The count map: &main.Counter{count:3}. [sender]
// The counter: &main.Counter{count:4}. [receiver]
// The count map: &main.Counter{count:4}. [sender]
// The counter: &main.Counter{count:5}. [receiver]
// The count map: &main.Counter{count:5}. [sender]
// Stopped. [receiver]
// $
