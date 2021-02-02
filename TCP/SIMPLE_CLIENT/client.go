package main

import (
	"log"
	"net"
    "fmt"
)

const (
    IP = "192.168.1.46"
    PORT = "3333"
    PROTO = "tcp"
)

func main(){
    // resolv given addr
    resolv, err := net.ResolveTCPAddr("tcp", IP+":"+PORT);
    
    // connect to the given addr
    conn, err := net.DialTCP("tcp", nil , resolv);
    if err != nil{log.Fatal("Can't resolve nameserver", err);}

    writeBuffer(conn);
    readBuffer(conn);
}

func writeBuffer(conn *net.TCPConn){
    str := "1";
    // Write in Buffer
    _, err := conn.Write([]byte(str));
    if err != nil{log.Fatal("Error writing string", err);}
}

func readBuffer(conn *net.TCPConn){
    // Create buffer to hold data from the tcp conn
    data := make([]byte, 1024);
    // Read Buffer
    n, err := conn.Read(data);
    if err != nil {log.Fatal("Error while Reading Buffer", err);}
    // Write data on screen
    s := string(data[:n]);
    fmt.Print(s);
}
