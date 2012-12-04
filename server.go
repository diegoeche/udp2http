package main

import(
        "fmt"
        "net"
        "bufio"
       )

func main () {
  ln, err := net.Listen("tcp", ":2112")
  if err != nil {
    fmt.Println("Listening Error!");
  }
  for {
      conn, err := ln.Accept()
      if err != nil {
        fmt.Println("Connection Error!");
	continue
      }
      go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  line, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Printf("Testing: %s", line)
}