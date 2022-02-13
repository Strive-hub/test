package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func hello(){
	defer wg.Done()
	fmt.Println("hello Goroutine!")
	//wg.Done()
}
func main(){
	wg.Add(1)
	go hello()
	fmt.Println("main goroutine done")
	wg.Wait()
}