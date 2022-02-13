package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	var host string
	flag.StringVar(&host,"host","host","curl host")
	flag.Parse()
	resp,err := http.Get(host)
	if err != nil{
		fmt.Println("http get failed ,err:",err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("read from resp.body failed,err:",err)
	}
	fmt.Println(string(body))
}
