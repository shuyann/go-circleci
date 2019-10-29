package main

import "testing"

func TestEcho(t *testing.T) {
	if echo() != "bar" {
		t.Errorf("echo() = %s but want %s", echo(), "bar")
	}
}
