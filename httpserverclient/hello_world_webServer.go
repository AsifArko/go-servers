package httpserverclient

import (
	"net/http"
	"fmt"
	"strings"
)

func HelloServer(w http.ResponseWriter , req *http.Request)  {
	fmt.Println("Inside Hello Server handler : ")
	fmt.Fprintf(w,"hello , "+req.URL.Path[1:])
}

func AboutPage(w http.ResponseWriter , req *http.Request)  {
	title := "About us"
	body  := "The fmt.Println statement prints on the server console; somewhat more useful could be to loginside every handler what was requested."
	fmt.Println("Inside About Page .")
	//fmt.Fprintf(w,"This is the about page , "+req.URL.Path[1:])
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, body)
	fmt.Fprintf(w,strings.Split(req.URL.Path,"/")[len(strings.Split(req.URL.Path,"/"))-1])
}

func main()  {
	http.HandleFunc("/",HelloServer)
	http.HandleFunc("/abouts",AboutPage)
	err := http.ListenAndServe("localhost:9999",nil)

	if err != nil{
		fmt.Println(err.Error())
	}
}