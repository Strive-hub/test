//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main()  {
//	output1 := make(chan string,10)
//	go write(output1)
//	for s,ok := range output1{
//		if ok{
//			fmt.Println("res:",s,"\n")
//			time.Sleep(time.Second)
//		}
//	}
//	close(output1)
//}
//
//func write(ch chan string) {
//	for i:=0;i<15;i++{
//		select {
//		case ch <- "hello":
//			fmt.Println("write hello")
//		default:
//			fmt.Println("channel full")
//		}
//		time.Sleep(time.Millisecond*500)
//	}
//}
//package main
//var c = make(chan int)
//var a int
//
//func f() {
//	a = 1
//	<-c
//}
//func main() {
//	print(&a)
//	go f()
//	print(&a)
//	c <- 0
//	print(&a)
//}