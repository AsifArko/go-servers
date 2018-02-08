package main

import (
	"net/http"
	"fmt"

	"reflect"
	"strings"
)

func Homepage( w http.ResponseWriter , req *http.Request)  {
	title := "Untitled"
	body := "This is the body"
	fmt.Fprintf(w,"<h1>%s</h1><div>%s</div>", title, body)
	fmt.Fprint(w,strings.Split(req.URL.Path,"/"))
}

func Exercise(w http.ResponseWriter , req *http.Request)  {
	title := "Exercise"
	body := "This is the body"
	fmt.Fprintf(w,"<h1>%s</h1><div>%s</div>", title, body)


	fmt.Println(req.URL.Path , reflect.TypeOf(req.URL.Path))

}

func main()  {
	exurl := "/ex"
	http.HandleFunc("/",Homepage)
	http.HandleFunc(exurl,Exercise)

	err := http.ListenAndServe("localhost:8080",nil)

	if err != nil{
		fmt.Println(err.Error())
	}


}
