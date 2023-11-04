package main

import (
        "log"
        "net"
        // "os"
)

func main() {
        ln, err := net.Listen("tcp", ":6000") 
        if err != nil {
                log.Fatal(err)
        }

        for {

                conn, err := ln.Accept()
                if err != nil {
                        log.Println(err)
                        continue
                }
                go handleConnection(conn)
        }
}

func handleConnection(c net.Conn) {
        log.Println("A client has connected", c.RemoteAddr())
        // c.Write([]byte(os.Getenv("hostname") ))
        c.Write([]byte("hkhkhkhh" ))
}
