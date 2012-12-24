package main

import(
        "fmt"
"net"
"bytes"
"encoding/binary"

)

// func binaryToHttp (b []byte) string{

func readString (buffer *bytes.Buffer) (string) {
	var stringLength int16
	binary.Read(buffer, binary.BigEndian, &stringLength)

	if stringLength > 1000 {
		stringLength = 1000
	}

	var str = make([]byte, stringLength)
	binary.Read(buffer, binary.BigEndian, &str)

	return string(str)
}

func binaryToHttp (b []byte) {
	buf := bytes.NewBuffer(b)
	var method	int16

	binary.Read(buf, binary.BigEndian, &method)

	var path = readString(buf)
	var body = readString(buf)

	fmt.Println("Method: ", method)
	fmt.Println("Path: ", string(path))
	fmt.Println("Body: ", string(body))
}

func main () {
	conn, err := net.ListenPacket("udp", ":2112")
	if err != nil {
		fmt.Println("Listening Error!");
	}
	var buf [1000]byte;
	for {
		conn.ReadFrom(buf[0:999]);
		binaryToHttp(buf[0:999])
	}
}

