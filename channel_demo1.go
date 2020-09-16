package main

import (
	"fmt"
)

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine打印", i)
			//time.Sleep(3*time.Second)
		}
		ch1 <- true
		fmt.Println("子goroutine结束")
	}()

	data := <-ch1
	fmt.Println("main...data-->", data)
	fmt.Println("main...over...")
}
