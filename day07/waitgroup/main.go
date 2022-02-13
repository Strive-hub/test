//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var (
//	x int
//	wg sync.WaitGroup
//	lock sync.Mutex
//	rwlock sync.RWMutex
//)
//func write(){
//	rwlock.Lock()
//	x +=1
//	time.Sleep(time.Millisecond*10)
//	rwlock.Unlock()
//	wg.Done()
//}
//func read(){
//	rwlock.RLock()
//	time.Sleep(time.Millisecond)
//	rwlock.RUnlock()
//	wg.Done()
//}
//func main(){
//	start := time.Now()
//	for i := 0;i<10;i++{
//		wg.Add(1)
//		go write()
//	}
//	for i := 0;i<10000;i++{
//		wg.Add(1)
//		go read()
//	}
//	wg.Wait()
//	end := time.Now()
//	fmt.Println(end.Sub(start))
//}