package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	file,err := os.Open("test.txt")
	if err !=nil{
		fmt.Printf("file '%v' open failed\n",file.Name())
	}else{
		fmt.Printf("file '%v' open success\n",file.Name())
	}
	defer file.Close()
	var buf [128]byte
	var content []byte
	for {
		n,err:= file.Read(buf[:])
		if err == io.EOF{
			break
		}
		if err != nil {
			fmt.Println("read file err",err)
			return
		}
		content = append(content,buf[:n]...)
	}
	fmt.Println(string(content))
}