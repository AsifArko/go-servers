package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func checkError(err error , str string){
	if err !=nil{
		fmt.Println(err.Error(),str)
		os.Exit(1)
	}

}

func main()  {
	url := "http://www.google.com"
	x:=os.Args

	fmt.Println(x)
	resp , err := http.Get(url)
	checkError(err,"Get Error")

	data , err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

}