package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
	fmt.Println("Failed to bind to port 4221")
	os.Exit(1)
	}
	
	defer l.Close()
	fmt.Println("Server is running on port 4221...")

	for{
		con, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(con)
	}
}

func handleConnection(con net.Conn){
	defer con.Close()

	// Read the incomming request
	buf := make([]byte,1024)
	n,err := con.Read(buf)
	if err != nil{
		fmt.Println("Error reading request:", err.Error())
		return
	}
	request := string(buf[:n])
	fmt.Println("Request recieved: \n",request)

	//Checking if it is valid
	if strings.Contains(request,"GET / ") {
		response := "HTTP/1.1 200 OK\r\n\r\n"
		con.Write([]byte(response))
	}else{
		response := "HTTP/1.1 404 Not Found\r\n\r\n"
		con.Write([]byte(response))
	}
}