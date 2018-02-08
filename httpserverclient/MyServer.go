package httpserverclient

import (
	"net/http"
	"fmt"
	"io/ioutil"
)


func Handler(w http.ResponseWriter , req *http.Request)  {

	html , err := ioutil.ReadFile("index.html")

	if err != nil{
		fmt.Println(err.Error())
	}

	fmt.Fprintf(w,string(html))

}

func main()  {
	http.HandleFunc("/",Handler)

	err := http.ListenAndServe("localhost:5432",nil)

	if err != nil{
		fmt.Println(err.Error())
	}


}
