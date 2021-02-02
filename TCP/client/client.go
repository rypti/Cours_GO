package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
    "log"
)

func main(){
    fmt.Println("Hello Friend :3");
    var network string 
   
    // Test if an argument is passed to the program
    if len(os.Args) <= 1 {
        log.Fatal("Please add a name to resolve")
    }else{ network = os.Args[1]}

    // Resolve the Domaine Name
    resolv, err := net.ResolveTCPAddr("tcp", network)
    if err != nil{log.Fatal("Can't resolve nameserver", err)}    

    fmt.Println("Connecting to : ", network, " -> ", resolv)
    // Starting the connection 
    conn, err  := net.DialTCP("tcp", nil, resolv)
    if err != nil{log.Fatal("Error Connection", err)}    

    // Starting the Reader in a go routine (Thread)
    go readBuffer(conn)
    // Reading until program dies
    go writeInBuffer(conn)
    for {}
}

func writeInBuffer(conn *net.TCPConn){
    for {
        // Create new Reader with STDIN as input
        str := bufio.NewReader(os.Stdin)
        // Read Until ENTER is pressed
        text, err := str.ReadString('\n')
        if err != nil {log.Fatal("Error reading string from STDIN", err)}
        _, err = conn.Write([]byte(text))
        if err != nil {log.Fatal("Error writing string", err)}
    }
}

func readBuffer(conn *net.TCPConn){
    for {
        data := make([]byte, 1024)
        n, err := conn.Read(data)
        if err != nil {log.Fatal("Error while Reading Buffer", err)}

        s := string(data[:n])
        fmt.Print(s)
    }
}
