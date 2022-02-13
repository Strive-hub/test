package main

import (
	"fmt"
	"reflect"
)

func main(){
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
		Test int `haha:"123"`
	}
	ins := cat{"mimi",1,2}
	typeOfCat := reflect.TypeOf(ins)
	for i :=0;i<typeOfCat.NumField();i++{
		fileType := typeOfCat.Field(i)
		fmt.Printf("name: %v tag: '%v'\n",fileType.Name,fileType.Tag)
	}
	if catType,ok := typeOfCat.FieldByName("Type");ok {
		fmt.Println(catType.Tag.Get("json"),catType.Tag.Get("id"))
	}
	fmt.Println("------------------------")
	fmt.Println(typeOfCat.Field(2).Tag.Get("haha"))
	fmt.Println(typeOfCat.NumField())
}