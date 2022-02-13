//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	output1 := make(chan string)
//	output2 := make(chan string)
//	go test1(output1)
//	go test2(output2)
//	select {
//	case s1 := <-output1:
//		fmt.Println("s1=", s1)
//	case s2 := <-output2:
//		fmt.Println("s2=", s2)
//	}
//}
//
//func test1(ch chan string) {
//	time.Sleep(time.Second * 5)
//	ch <- "test1"
//}
//func test2(ch chan string) {
//	time.Sleep(time.Second * 2)
//	ch <- "test2"
//}
