package main

import (
    "fmt"
    "log"
    "net"
    "time"
)

func main() 
{

    // listen on a port
    ln, err := net.Listen("tcp", ":8080")
    defer ln.Close()
    if err != nil 
	{
        log.Panic(err)
    }

    for {
        // accept a connection
        c, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
            continue
        }

        // handle the connection
        go func(conn net.Conn) 
		{

            daytime := time.Now().String() + " ...Hallo\n"
            // conn.Write([]byte(daytime))
            fmt.Fprintf(conn, daytime)
            conn.Close()
        }(c)
    }
}