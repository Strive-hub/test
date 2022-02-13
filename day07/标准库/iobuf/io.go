package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wr(){
	file,err := os.OpenFile("./test.txt",os.O_CREATE|os.O_WRONLY,0666)
	if err !=nil{
		fmt.Println("file create failed",err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i:=0;i<10;i++{
		writer.WriteString("hello\n")
	}
	writer.Flush()
}
func re(){
	file,err := os.Open("./test.txt")
	if err != nil{
		fmt.Println("file open failed ",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line,_,err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		if err != nil{
			return
		}
		fmt.Println(string(line))
	}
}
func main()  {
	wr()
	re()
}