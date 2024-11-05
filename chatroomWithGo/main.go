package main

import "net"

func main() {
	println("This is our chatroom")
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		println(err)
	}

	// Accepting connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			println(err)
		}
		// Create a go routine to handleConnection
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	// Close the connection when we're done
	// defer is a keyword that schedules a function to be executed after the surounding code is finished, whether by error or just finishe

	defer conn.Close()
	// Read the incoming data
	buf := make([]byte, 1024)
	for {

		_, err := conn.Read(buf)

		if err != nil {
			println(err)
			return
		}

		println("Received Data : " + string(buf))
	}
}
