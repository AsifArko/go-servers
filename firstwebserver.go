package main

import (
	"net/http"
	"io"
)

const form = `<html>
			<body>
			<form 
				action="#" method="post" name="bar"><input type="text" name="in"/><input type="submit" value="Submit"/>
			</form>
			</body>
			</html>`

func Server(w http.ResponseWriter , request *http.Request)  {
	io.WriteString(w,"<h1>Strange Things</h2>")
}

func FServer(w http.ResponseWriter , request *http.Request)  {
	w.Header().Set("Content-Type","text/html")

	switch request.Method {
	case "GET":
		io.WriteString(w,form)
	case "POST":
		io.WriteString(w,request.FormValue("in"))
	}
}

func main()  {
	http.HandleFunc("/",Server)
	http.HandleFunc("/strangers",FServer)

	if err := http.ListenAndServe(":8091",nil);err!=nil{
		panic(err)
	}
}