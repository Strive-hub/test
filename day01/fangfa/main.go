package main

import "fmt"

type Sayer interface {
	say()
}
type dog struct {}
type cat struct {}
func (d dog) say(){
	fmt.Println("dog")
}
func (c cat) say(){
	fmt.Println("cat")
}
func main(){
	var x Sayer
	a := new(cat)
	b := dog{}
	x = a
	x.say()
	x = b
	x.say()
}