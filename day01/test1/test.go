package main

import "fmt"

type Bag struct {
	itmes []int
}
func Insert(b *Bag,itemid int){
	b.itmes = append(b.itmes,itemid)
}
func main(){
	bag := new(Bag)
	Insert(bag,102)
	fmt.Println(bag)
}