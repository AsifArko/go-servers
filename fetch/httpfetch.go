package fetch

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func checkErr(err error , string string)  {
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main()  {
	res , err := http.Get("http://google.com")

	checkErr(err , "Http Get error")

	data , err := ioutil.ReadAll(res.Body)
	checkErr(err , "Read response")

	//fmt.Printf("%q",string(data))
	fmt.Println(string(data))
}
