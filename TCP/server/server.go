package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
    var network string; 
    // Test if an argument is passed to the program
    if len(os.Args) <= 1 {
        log.Fatal("Please add an IP and a port to listen on ");
    }else {network = os.Args[1];}
    
    // listen for incoming connection
    lis, err := net.Listen("tcp", network);
    if err != nil{log.Fatal("Can't listen on given ip and port ", err);}

    defer lis.Close();

    fmt.Println("Listening on " + network);
    
    for {
        // Listen for an incoming connection.
        conn, err := lis.Accept();
        if err != nil{log.Fatal("Error accepting ", err);}
       
        // Handle connections in a new goroutine.
        go handleRequest(conn);
    }
}

func handleRequest(conn net.Conn) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 2048);

    // Read incoming connection
    reqLen, err := conn.Read(buf);
    if err != nil {fmt.Println("Error reading ", err);}
   
    // Buffer to string
    s := string(buf[:reqLen]);
    fmt.Println(conn.RemoteAddr(),":",s);

    // Write buffer to Connection
    conn.Write([]byte("Message received.\n"));

    // Close the connection
    conn.Close();
}

