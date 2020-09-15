package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//并发互斥锁模型

var tickers = 10 // 十张车票
var wg sync.WaitGroup // 同步等待组
var mutex sync.Mutex // 互斥锁


func main(){
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")
	wg.Wait()
}

func saleTickets(name string){
	defer wg.Done() // 同步锁执行完毕
	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock() //上锁需要加在判断之前
		if tickers > 0 {
			//休息一下 更方便展示效果
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出:", tickers)
			tickers --
		} else {
			mutex.Unlock()
			fmt.Println(name, "售罄！")
			break
		}
		mutex.Unlock()
	}
}