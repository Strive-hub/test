package main

import "fmt"

var x = 10
func main(){
	 fmt.Println(x)
	x :=1
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
}
