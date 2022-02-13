package main

import (
	"fmt"
	"reflect"
)

type Myint int
type User struct {
	name string
	age  int
}

func main() {
	u := User{"lyd", 24}
	var num Myint = 18
	v := reflect.ValueOf(u)
	vNum := reflect.ValueOf(num)
	y := v.Interface().(User)
	//yNum := vNum.Interface()
	yNum := vNum.Interface()
	fmt.Println(y, yNum)
}
