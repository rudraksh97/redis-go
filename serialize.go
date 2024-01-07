package main

import "fmt"

func serializeInteger(i int) string {
	return ":" + fmt.Sprint(i) + "\r\n"
}

func serializeSimpleString(s string) string {
	return "+" + s + "\r\n"
}

func serializeBulkString(str string) string {
	return "$" + fmt.Sprint(len(str)) + "\r\n" + str + "\r\n"
}

func serializeError(s string) string {
	return "-" + s + "\r\n"
}

func serializeArray(arr []string) string {
	result := fmt.Sprintf("*%d\r\n", len(arr))
	for _, item := range arr {
		result += serializeBulkString(item)
	}
	return result
}

func serializeNull() string {
	return "$-1\r\n"
}
