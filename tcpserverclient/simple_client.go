package tcpserverclient

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

const address = "localhost:4444"
func main()  {

	conn , err := net.Dial("tcp",address)

	if err != nil{
		fmt.Println(err.Error(),"Error dialing .")
		return
	}


	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name : ")
	clientName , _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName,"\n")

	for{
		fmt.Printf("%s > ",trimmedClient)
	input , _ := inputReader.ReadString('\n')
	trimmedInput := strings.Trim(input,"\n")

	if trimmedInput == "QUIT"{
		return
	}

	_, err := conn.Write([]byte(trimmedClient+" : "+trimmedInput))

	if err!=nil{
		fmt.Print(err.Error(),"Write Error .")
	}

	}



}
