//package main
//
//import "fmt"
//
//func main()  {
//	int_chan := make(chan int,1)
//	string_chan := make(chan string,1)
//	go func() {
//		int_chan <- 1
//	}()
//	go func() {
//		string_chan <- "hello"
//	}()
//	select {
//	case value := <- int_chan:
//		fmt.Println(value)
//	case value := <- string_chan:
//		fmt.Println(value)
//	}
//	fmt.Println("main end")
//}
