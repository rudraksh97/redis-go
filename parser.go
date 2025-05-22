package main

import (
	"bytes"
	"errors"
	"strings"
)

// ParseCommand parses a raw byte slice into a command and its arguments.
// It assumes commands are simple space-separated strings on a single line,
// ending with CRLF.
// e.g., "SET key value\r\n" -> ["SET", "key", "value"], nil
func ParseCommand(data []byte) ([]string, error) {
	// Trim CRLF and any other leading/trailing whitespace
	trimmedData := bytes.TrimSpace(data)

	if len(trimmedData) == 0 {
		return nil, errors.New("ERR empty command")
	}

	// Split into command and arguments
	// Using Fields to handle multiple spaces between arguments correctly
	parts := strings.Fields(string(trimmedData))

	if len(parts) == 0 {
		// This case should ideally not be reached if len(trimmedData) > 0,
		// but as a safeguard.
		return nil, errors.New("ERR empty command after parsing")
	}

	return parts, nil
}
