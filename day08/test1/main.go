package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	//time := time.Now()
	//logname := fmt.Sprintf("%d-%02d-%02d",time.Year(),time.Month(),time.Day())
	//fmt.Println(logname)
	logFile,err := os.OpenFile("./test.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil{
		fmt.Println("open log file failed,err:",err)
		return
	}
	log.SetFlags(log.Lshortfile|log.Lmicroseconds|log.Ldate)
	log.SetOutput(logFile)
	log.Println("This is a nomor log")
	log.SetPrefix("[pprof]")
	log.Println("This is a nomor log")
}