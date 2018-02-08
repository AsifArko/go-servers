package tcpserverclient

import (
	"flag"
	"fmt"
	"net"
)

const maxRead = 25


func main()  {
	flag.Parse()

	if flag.NArg() != 2{
		panic("usage:host:port")
	}

	hostAndPort := fmt.Sprintf("%s:%s",flag.Arg(0),flag.Arg(1))
	listener := initServer(hostAndPort)

	for {
		conn , err := listener.Accept()
		checkError(err,"Accept : ")
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string)(*net.TCPListener)  {
	serverAddr , err := net.ResolveTCPAddr("tcp",hostAndPort)
	checkError(err,"Resolving address : port failed"+ hostAndPort+"")
	listener , err := net.ListenTCP("tcp",serverAddr)
	checkError(err,"ListenTCP: ")
	fmt.Println("Listening to : ",listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn){
	connFrom := conn.RemoteAddr().String()
	fmt.Println("connection from : ",connFrom)
	sayHallo(conn)
	for{
		var ibuf []byte = make([]byte,maxRead+1)
		length,err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil :
			HandleMsg(length,err,ibuf)
		default:
			goto DISCONNECT

		}
	}

DISCONNECT:
	err := conn.Close()
	fmt.Println("Closed connection : ",connFrom)
	checkError(err,"Close")
}

func sayHallo(to net.Conn)  {
	obuf := []byte{'l','e','t','s'}
	wrote , err := to.Write(obuf)
	checkError(err,"Write : wrote "+string(wrote)+" bytes.")
}

func GetUser(msg []byte) string {
	var user string
	for i:=0;;i++{
		if msg[i]==58{
			break
		}else {
			user = user+fmt.Sprintf("%c",msg[i])
		}
	}
	return user
}

func HandleMsg(length int , err error , msg []byte) string{
	var user []string

	if length > 0 {
		for i:=0;;i++{
			if msg[i] == 0{
				break
			}
			fmt.Printf("%c",msg[i])
		}
		fmt.Println()
	}
	userName := GetUser(msg)
	user = append(user, userName)
	return "string"
}

func checkError(error error , info string)  {
	if error != nil{
		panic("Error : "+info+""+error.Error())
	}
}
