package main

import (
	"net/http"
	"io"
	"fmt"
	"os"
	"bufio"
	"strconv"
)


func SimpleServer(w http.ResponseWriter , request *http.Request)  {
	io.WriteString(w,"<h1>Hello World</h1>")
}

func FormServer(w http.ResponseWriter , request *http.Request)  {
	form := GetHtml()
	w.Header().Set("Content-Type","text/html")

	switch request.Method{
	case "GET":
		io.WriteString(w,form)
	case "POST":
		x := request.FormValue("firstname")
		io.WriteString(w,request.FormValue("firstname"))
		io.WriteString(w,request.FormValue("lastname"))

		fmt.Println(strconv.Atoi(x))
	}
}

func main()  {

	http.HandleFunc("/test1",SimpleServer)
	http.HandleFunc("/test2",FormServer)

	if err := http.ListenAndServe(":8088",nil);err!=nil{
		panic(err)
	}


}


func GetHtml() string  {

	var data string

	x,err:=os.Open("/home/hoods/go/src/github.com/gonetworking/simplewebserver/page.html")

	if err != nil{
		fmt.Println(err.Error())
	}

	ir :=bufio.NewReader(x)

	for{
		line ,err := ir.ReadString('\n')

		if err == io.EOF{
			break
		}
		data = data+line
	}
	return data
}
