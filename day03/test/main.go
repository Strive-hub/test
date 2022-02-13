package main

import "fmt"

func main()  {
	slice := []int{1,2,3,4,5,6}
	for range slice{
		fmt.Println(slice)
	}
}
