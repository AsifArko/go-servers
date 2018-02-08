package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main()  {

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
		fmt.Print(line)
	}
}
