package main

import (
	"fmt"
	"sync"
	"time"
)

func lock()  {
	lock := sync.Mutex{}
	for i :=0;i<100;i++{
		lock.Lock()
		defer lock.Unlock()
		fmt.Println(i)
		//lock.Unlock()
	}
}
func rlock(){
	lock := sync.RWMutex{}
	for i :=0;i<10;i++{
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println(i)
	}
}
func wlock(){
	lock := sync.RWMutex{}
	for i:=0;i<100;i++{
		//lock.Lock()
		//defer lock.Unlock()
		go func(i int) {
			lock.Lock()
			//defer lock.Unlock()
			fmt.Println(i)
			lock.Unlock()
		}(i)
		//lock.Unlock()
	}
}
func main()  {
	//go lock()
	//go rlock()
	//time.Sleep(time.Second)
	go wlock()
	time.Sleep(time.Second*10)
}
