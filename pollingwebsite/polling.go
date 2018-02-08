package main

import (
	"net/http"
	"fmt"
)

var urls = []string{
	"http://google.com",
	"http://golang.org",
	"http://blog.golang.org",
}

func main()  {
	for _, url := range urls{
		resp , err := http.Head(url)

		if err != nil{
			fmt.Println(err.Error())
		}

		fmt.Println(url , " : " , resp.Request , resp.Header)
	}
}