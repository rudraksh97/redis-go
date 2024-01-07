package main

import (
	"fmt"
	"strconv"
	"strings"
)

func deserialize(s string) (any, error) {
	const INDEX = 0

	if s[INDEX] == ':' {
		return deserializeInteger(s)
	} else if s[INDEX] == '+' {
		return deserializeSimpleString(s)
	} else if s[INDEX] == '-' {
		return deserializeError(s)
	} else if s[INDEX] == '*' {
		return deserializeArray(s)
	} else if s[INDEX] == '$' {
		fmt.Println(s[INDEX+1])
		if s[INDEX+1] == '-' {
			return nil, nil
		}
		return deserializeBulkString(s)
	} else {
		return nil, fmt.Errorf("invalid prefix in serialised message")
	}
}

func deserializeInteger(s string) (string, error) {
	return s[1 : len(s)-2], nil
}

func deserializeSimpleString(s string) (string, error) {
	return s[1 : len(s)-2], nil
}

func deserializeError(s string) (string, error) {
	return s[1 : len(s)-2], nil
}

func deserializeArray(s string) ([]string, error) {
	parts := strings.Split(s, "\r\n")

	sizeOfArray, err := strconv.Atoi(parts[0][1:])

	if err != nil {
		return nil, fmt.Errorf("invalid count in RESP array")
	}

	elements := make([]string, 0, sizeOfArray)
	for i := 1; i < len(parts) && len(elements) < sizeOfArray; i += 2 {
		if strings.HasPrefix(parts[i], "$") {
			elements = append(elements, parts[i+1])
		}
	}

	if len(elements) != sizeOfArray {
		return nil, fmt.Errorf("mismatched count and number of elements in RESP array")
	}

	return elements, nil
}

func deserializeBulkString(s string) (string, error) {
	parts := strings.Split(s, "\r\n")

	sizeOfString, err := strconv.Atoi(parts[0][1:])

	if err != nil {
		return "", fmt.Errorf("invalid in RESP bulk string")
	}

	element := parts[1]

	if len(element) != sizeOfString {
		return "", fmt.Errorf("mismatched count and number of bytes in RESP bulk string")
	} else {
		return element, nil
	}
}
