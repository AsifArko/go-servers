package main

import (

	"net/http"
	"fmt"
	"io/ioutil"

)

type Status struct {
	Text string
}

type User struct{
	Name string
	Status Status
}

func main()  {
	//var b = []byte{}
	resp , _ := http.Get("https://twitter.com/arko_asif/status/961284130888495105")

	data , err := ioutil.ReadAll(resp.Body)

	if err != nil{
		fmt.Println(err.Error())
	}

	fmt.Println(string(data))
}
