package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		deserializedMessage, _ := readAndDeserialize(reader)
		switch deserializedMessage {
		case "PING":
			conn.Write([]byte(serializeSimpleString("PONG")))
		case "ECHO":
			conn.Write([]byte(serializeSimpleString("Tujha Aaichi Gaand")))
		default:
			conn.Write([]byte(serializeError("ERR unknown command")))
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":6379")

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()

	fmt.Println("Listening on :6379 for connections...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}

		go handleConnection(conn)
	}
}

func readAndDeserialize(reader *bufio.Reader) (any, error) {
	message, err := getMessage(reader)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return nil, err
	}

	deserializedMessage, err := deserialize(message)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return nil, err
	}

	return deserializedMessage, nil
}

// func getCompleteMessage(reader *bufio.Reader) (string, error) {
// 	message, err := getMessage(reader)
// }

func getMessage(reader *bufio.Reader) (string, error) {
	message, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return "", err
	}

	fmt.Printf("Command received: %s", message)

	trimmedMessage := strings.TrimSpace(message)
	return trimmedMessage, nil
}
