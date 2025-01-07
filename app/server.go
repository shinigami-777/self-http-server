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

	var useragent string
	useragent = "nil"
	var host string
	host = "nil"

	// Header Parsing here
	lines := strings.Split(request, "\r\n")
	for _, line := range lines {
		fmt.Println(line)
		if strings.HasPrefix(line, "User-Agent:") {
			useragent = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		}
		if strings.HasPrefix(line, "Host:") {
			host = strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
		}
	}
	fmt.Println(host)   //Only did this for that declared and not used error

	fmt.Println(lines)
	requestLine := strings.Fields(lines[0])
	method := requestLine[0]
	path := requestLine[1]
	//version :=requestLine[2]
	//requestLine = strings.Fields(lines[1])
	//host := requestLine[1]
	//fmt.Println(version)
	//fmt.Println(host)
	//requestLine = strings.Fields(lines[2])
	//useragent := requestLine[1]
	//fmt.Println(useragent)
	//fmt.Println(requestLine)  //to check the lines
	

	if method == "GET" {
		if strings.HasPrefix(path, "/echo/"){
			str := strings.TrimPrefix(path, "/echo/")
			contentLength := len(str)
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",contentLength, str)
			con.Write([]byte(response))
			fmt.Println(response)
			return;
		}
		if path == "/user-agent" || path == "/user-agent/"{
			contentLength := len(useragent)
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",contentLength, useragent)
			con.Write([]byte(response))
			fmt.Println(response)
			return;
		}
	}

	//Checking if it is valid
	if strings.Contains(request,"GET / ") {
		response := "HTTP/1.1 200 OK\r\n\r\n"
		con.Write([]byte(response))
	}else{
		response := "HTTP/1.1 404 Not Found\r\n\r\n"
		con.Write([]byte(response))
	}
}