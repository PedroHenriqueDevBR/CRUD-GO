package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	name := "server"
	if name != "server" {
		t.Fatalf("incorrect name")
	}
}
