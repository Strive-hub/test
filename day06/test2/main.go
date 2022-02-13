package main

import (
	"fmt"
	"reflect"
)

func main(){
	type myint int
	num := myint(2)
	fmt.Println(num)
	fmt.Println(reflect.TypeOf(num),reflect.ValueOf(num).Kind())
}
