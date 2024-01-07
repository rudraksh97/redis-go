package main

import (
	"reflect"
	"testing"
)

func TestDeserializationNull(t *testing.T) {
	msg, err := deserialize("$-1\r\n")
	if msg != nil || err != nil {
		t.Fatalf("Failed")
	}
}

func TestDeserializationArray(t *testing.T) {
	msg, err := deserialize("*1\r\n$4\r\nping\r\n")
	want := []string{"ping"}
	if !reflect.DeepEqual(msg, want) || err != nil {
		t.Fatalf("Failed")
	}
}

func TestDeserializationArray2(t *testing.T) {
	msg, err := deserialize("*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n")
	want := []string{"echo", "hello world"}
	if !reflect.DeepEqual(msg, want) || err != nil {
		t.Fatalf("Failed")
	}
}

func TestDeserializationArray3(t *testing.T) {
	msg, err := deserialize("*2\r\n$3\r\nget\r\n$3\r\nkey\r\n")
	want := []string{"get", "key"}
	if !reflect.DeepEqual(msg, want) || err != nil {
		t.Fatalf("Failed")
	}
}

func TestDeserializationArray4(t *testing.T) {
	msg, err := deserialize("*2\r\n$3\r\nget\r\n$3\r\nkey\r\n")
	want := []string{"get", "key"}
	if !reflect.DeepEqual(msg, want) || err != nil {
		t.Fatalf("Failed")
	}
}

func TestDeserializationSimpleString(t *testing.T) {
	msg, err := deserialize("+OK\r\n")
	want := "OK"
	if msg != want || err != nil {
		t.Fatalf("Failed")
	}
}

func TestDeserializationError(t *testing.T) {
	msg, err := deserialize("-Error message\r\n")
	want := "Error message"
	if msg != want || err != nil {
		t.Fatalf("Failed")
	}
}

func TestDeserializationBulkString(t *testing.T) {
	msg, err := deserialize("$0\r\n\r\n")
	want := ""
	if msg != want || err != nil {
		t.Fatalf("Failed")
	}
}

func TestDeserializationSimpleString2(t *testing.T) {
	msg, err := deserialize("+hello world\r\n")
	want := "hello world"
	if msg != want || err != nil {
		t.Fatalf("Failed")
	}
}
