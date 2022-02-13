package main

import (
	"fmt"
	"sync"
	"time"
)
var cond *sync.Cond
func init(){
	cond = sync.NewCond(&sync.Mutex{})
}
func test(x int){
	cond.L.Lock()
	cond.Wait()
	fmt.Println(x)
	time.Sleep(time.Second)
	cond.L.Unlock()
}
func main(){
	cond := sync.NewCond(&sync.Mutex{})
	for i :=0;i<40;i++{
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second*2)
	fmt.Println("broadcast")
	cond.Signal()
	time.Sleep(time.Second*2)
	cond.Signal()
	time.Sleep(time.Second*2)
	cond.Broadcast()
}