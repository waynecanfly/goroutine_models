// 读写锁并发模型
package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)
	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("main over...")

}

func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始写。。")
	rwMutex.Lock() //写上锁
	fmt.Println(i, "正在写。。")
	time.Sleep(3 * time.Second)
	fmt.Println(i, "写结束")
	rwMutex.Unlock() //写结束 解锁
}

func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始读。。")
	rwMutex.RLock() //读上锁
	fmt.Println(i, "正在读。。")
	time.Sleep(3 * time.Second)
	fmt.Println(i, "读结束。。")
	rwMutex.RUnlock() //读结束， 释放读锁
}
