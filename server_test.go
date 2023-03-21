package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	name := "serverino"
	if name != "server" {
		t.Fatalf("incorrect name")
	}
}
