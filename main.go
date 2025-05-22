package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"strings"
)

// Global store instance
var store = NewStore()

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accepted new connection from %s\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)

	for {
		// Read up to the \n delimiter
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Client %s disconnected.\n", conn.RemoteAddr())
			} else {
				fmt.Printf("Error reading from client %s: %s\n", conn.RemoteAddr(), err.Error())
			}
			return // Close connection on EOF or error
		}

		// Parse the command
		parts, err := ParseCommand(line)
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("-ERR %s\r\n", err.Error())))
			continue
		}

		if len(parts) == 0 {
			// Should be caught by ParseCommand, but as a safeguard
			conn.Write([]byte("-ERR empty command\r\n"))
			continue
		}

		command := strings.ToUpper(parts[0])
		args := parts[1:]

		switch command {
		case "PING":
			if len(args) != 0 {
				conn.Write([]byte("-ERR wrong number of arguments for 'ping' command\r\n"))
			} else {
				conn.Write([]byte("+PONG\r\n"))
			}
		case "SET":
			if len(args) != 2 {
				conn.Write([]byte("-ERR wrong number of arguments for 'set' command\r\n"))
			} else {
				store.Set(args[0], args[1])
				conn.Write([]byte("+OK\r\n"))
			}
		case "GET":
			if len(args) != 1 {
				conn.Write([]byte("-ERR wrong number of arguments for 'get' command\r\n"))
			} else {
				value, ok := store.Get(args[0])
				if ok {
					conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)))
				} else {
					conn.Write([]byte("$-1\r\n")) // Null bulk string
				}
			}
		default:
			conn.Write([]byte(fmt.Sprintf("-ERR unknown command '%s'\r\n", command)))
		}
	}
}

func main() {
	port := flag.String("port", "6379", "Port to listen on")
	flag.Parse()

	// store is already initialized as a global variable
	// store := NewStore() // No longer needed here

	ln, err := net.Listen("tcp", ":"+*port)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()

	fmt.Printf("Listening on :%s for connections...\n", *port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return // Changed from continue to return to stop if accept fails critically
		}

		// Pass the store instance to the handler
		// Since store is global now, no need to pass it as argument.
		// go handleConnection(conn, store)
		go handleConnection(conn)
	}
}

// The following functions are no longer needed as parsing and command handling
// are now done directly in handleConnection using ParseCommand.
// func readAndDeserialize(reader *bufio.Reader) (any, error) {
// 	message, err := getMessage(reader)

// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 		return nil, err
// 	}

// 	deserializedMessage, err := deserialize(message)

// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 		return nil, err
// 	}

// 	return deserializedMessage, nil
// }

// func getMessage(reader *bufio.Reader) (string, error) {
// 	message, err := reader.ReadString('\n')

// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 		return "", err
// 	}

// 	fmt.Printf("Command received: %s", message)

// 	trimmedMessage := strings.TrimSpace(message)
// 	return trimmedMessage, nil
// }
